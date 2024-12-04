package dijkstra

import (
	"context"
	"fmt"
	"io"

	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/apiserver/pkg/admission"
	"k8s.io/klog"

	informers "jinli.io/shortestpath/generated/client/informers/externalversions"
	v1listers "jinli.io/shortestpath/generated/client/listers/dijkstra/v1"
	v2listers "jinli.io/shortestpath/generated/client/listers/dijkstra/v2"
	"jinli.io/shortestpath/pkg/admission/custominitializer"
	"jinli.io/shortestpath/pkg/apis/dijkstra"
)

func Register(plugins *admission.Plugins) {
	plugins.Register("Dijkstra", func(config io.Reader) (admission.Interface, error) {
		return New()
	})
}

type DijkstraPlugin struct {
	// 实现了admission.Interface，可以被注册
	*admission.Handler
	knListerV1 v1listers.KnownNodesLister
	dpListerV1 v1listers.DisplayLister
	knListerV2 v2listers.KnownNodesLister
	dpListerV2 v2listers.DisplayLister
}

func New() (*DijkstraPlugin, error) {
	return &DijkstraPlugin{
		Handler: admission.NewHandler(admission.Create, admission.Update),
	}, nil
}

// 实现自定义插件接口
var _ custominitializer.WantsInformerFactory = &DijkstraPlugin{}
var _ admission.InitializationValidator = &DijkstraPlugin{}
var _ admission.ValidationInterface = &DijkstraPlugin{}

func (dp *DijkstraPlugin) SetInformerFactory(f informers.SharedInformerFactory) {
	// admission.PluginInitializer自动调用，启动lister
	dp.knListerV1 = f.Dijkstra().V1().KnownNodeses().Lister()
	dp.SetReadyFunc(f.Dijkstra().V1().KnownNodeses().Informer().HasSynced)

	dp.dpListerV1 = f.Dijkstra().V1().Displays().Lister()
	dp.SetReadyFunc(f.Dijkstra().V1().Displays().Informer().HasSynced)

	dp.knListerV2 = f.Dijkstra().V2().KnownNodeses().Lister()
	dp.SetReadyFunc(f.Dijkstra().V2().KnownNodeses().Informer().HasSynced)

	dp.dpListerV2 = f.Dijkstra().V2().Displays().Lister()
	dp.SetReadyFunc(f.Dijkstra().V2().Displays().Informer().HasSynced)
}

func (dp *DijkstraPlugin) ValidateInitialization() error {
	// admission.InitializationValidator自动调用，初始化阶段验证lister
	if dp.knListerV1 == nil || dp.dpListerV1 == nil {
		return fmt.Errorf("missing lister v1")
	}

	if dp.knListerV2 == nil || dp.dpListerV2 == nil {
		return fmt.Errorf("missing lister v2")
	}

	return nil
}

func (dp *DijkstraPlugin) Validate(ctx context.Context, a admission.Attributes, _ admission.ObjectInterfaces) error {
	// admission.ValidationInterface自动调用
	klog.Info("777777777777777777")
	klog.Info(a.GetSubresource())
	klog.Info("999999999999999999999")
	// 不支持子资源Verbs[Create]
	if a.GetSubresource() != "" {
		if a.GetOperation() == admission.Create {
			admission.NewForbidden(a, fmt.Errorf("not support subresource create operation"))
		}
	}

	// 验证阶段验证APIServer可用性
	if !dp.WaitForReady() {
		return admission.NewForbidden(a, fmt.Errorf("API Server not yet ready to handle request"))
	}

	if a.GetKind().GroupKind() != dijkstra.Kind("KnownNodes") && a.GetKind().GroupKind() != dijkstra.Kind("Display") {
		return nil
	}

	// 限制kn总数
	kns, err := dp.knListerV2.List(labels.Everything())
	if err != nil {
		return err
	} else if len(kns)+1 > custominitializer.KN {
		return fmt.Errorf("the number of kns is greater than %d", custominitializer.KN)
	}
	// 限制dp总数
	dps, err := dp.dpListerV2.List(labels.Everything())
	if err != nil {
		return err
	} else if len(dps)+1 > custominitializer.DP {
		return fmt.Errorf("the number of dps is greater than %d", custominitializer.DP)
	}

	// 限制单个kn对应的dp数量不能超过对应nodes数量
	if a.GetKind().GroupKind() == dijkstra.Kind("KnownNodes") {
		obj := a.GetObject()
		knObj := obj.(*dijkstra.KnownNodes)
		labelSelector := labels.Set{
			"nodeIdentity": knObj.Spec.NodeIdentity,
		}.AsSelector()
		ds, err := dp.dpListerV2.Displays(knObj.Namespace).List(labelSelector)
		if err != nil {
			return err
		}
		// 更新kn对象时校验,即当你减少nodes，应该删除对应的dp
		if len(ds) > len(knObj.Spec.Nodes) {
			klog.Warningf("the number of dps is greater than kn nodes for %s/%s,should delete extra display", knObj.Namespace, knObj.Spec.NodeIdentity)
		}
	} else if a.GetKind().GroupKind() == dijkstra.Kind("Display") {
		obj := a.GetObject()
		dpObj := obj.(*dijkstra.Display)
		labelSelector := labels.Set{
			"nodeIdentity": dpObj.Spec.NodeIdentity,
		}.AsSelector()

		ks, err := dp.knListerV2.KnownNodeses(dpObj.Namespace).List(labelSelector)
		if err != nil || len(ks) == 0 {
			return err
		}
		// 创建dp对象时校验
		ds, err := dp.dpListerV2.Displays(dpObj.Namespace).List(labelSelector)
		if err != nil {
			return err
		}

		totalDp := 0
		// 添加dp对象时校验
		if a.GetOperation() == admission.Create {
			totalDp = len(ds) + 1
		} else {
			totalDp = len(ds)
		}
		if totalDp > len(ks[0].Spec.Nodes) {
			return admission.NewForbidden(a, fmt.Errorf("the number of dps is greater than kn nodes for %s/%s", dpObj.Namespace, dpObj.Spec.NodeIdentity))
		}
	}

	return nil
}
