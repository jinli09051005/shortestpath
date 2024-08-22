package options

import (
	clientset "jinli.io/shortestpath/generated/client/clientset/versioned"
	informers "jinli.io/shortestpath/generated/client/informers/externalversions"
	"jinli.io/shortestpath/pkg/admission/custominitializer"
	djplugin "jinli.io/shortestpath/pkg/admission/plugins/dijkstra"
	"jinli.io/shortestpath/pkg/apis/dijkstra"
	v1 "jinli.io/shortestpath/pkg/apis/dijkstra/v1"
	v2 "jinli.io/shortestpath/pkg/apis/dijkstra/v2"
	utilerrors "k8s.io/apimachinery/pkg/util/errors"
	"k8s.io/apiserver/pkg/admission"
	genericapiserver "k8s.io/apiserver/pkg/server"
	genericoptions "k8s.io/apiserver/pkg/server/options"
)

// 自定义资源在etcd中存储路径
const defaultEtcdPathPrefix = "/registry/dijkstra.jinli.io"

type Options struct {
	RecommendedOptions    *genericoptions.RecommendedOptions
	SharedInformerFactory informers.SharedInformerFactory
}

func CreateOptions() *Options {
	opts := &Options{
		RecommendedOptions: genericoptions.NewRecommendedOptions(
			defaultEtcdPathPrefix,
			// 使用LegacyCodec版本的编解码器（ LegacyCodec 可以确保你的应用程序能够与 Kubernetes API 的旧版本兼容）
			dijkstra.Codecs.LegacyCodec(v2.SchemeGroupVersion, v1.SchemeGroupVersion),
		),
	}

	return opts
}

// options验证方法
func (o *Options) Validate() error {
	// Validate verifies flags passed to AdmissionOptions.
	if errs := o.RecommendedOptions.Validate(); len(errs) != 0 {
		return utilerrors.NewAggregate(errs)
	}
	return nil
}

// options配置方法
func (o *Options) Complete() error {
	o.RecommendedOptions.ExtraAdmissionInitializers = func(c *genericapiserver.RecommendedConfig) ([]admission.PluginInitializer, error) {
		client, err := clientset.NewForConfig(c.LoopbackClientConfig)
		if err != nil {
			return nil, err
		}
		informerFactory := informers.NewSharedInformerFactory(client, c.LoopbackClientConfig.Timeout)
		o.SharedInformerFactory = informerFactory
		return []admission.PluginInitializer{custominitializer.New(informerFactory)}, nil
	}
	// register admission plugins to options
	djplugin.Register(o.RecommendedOptions.Admission.Plugins)
	// 追加到有序插件列表,名称与注册时相同
	o.RecommendedOptions.Admission.RecommendedPluginOrder = append(o.RecommendedOptions.Admission.RecommendedPluginOrder, "Dijkstra")
	return nil
}
