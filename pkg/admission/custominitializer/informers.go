package custominitializer

import (
	informers "jinli.io/shortestpath/generated/client/informers/externalversions"
	"k8s.io/apiserver/pkg/admission"
)

// 提供informers
type informerPluginInitializer struct {
	SharedInformerFactory informers.SharedInformerFactory
}

var _ admission.PluginInitializer = informerPluginInitializer{}

func (i informerPluginInitializer) Initialize(plugin admission.Interface) {
	// 如果插件实现了WantsInformerFactory接口，调用插件的SetInformerFactory方法并传递informerFactory启动lister
	if wants, ok := plugin.(WantsInformerFactory); ok {
		wants.SetInformerFactory(i.SharedInformerFactory)
	}
}

func New(informerFactory informers.SharedInformerFactory) informerPluginInitializer {
	return informerPluginInitializer{
		SharedInformerFactory: informerFactory,
	}
}
