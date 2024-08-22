package install

import (
	"jinli.io/shortestpath/pkg/apis/dijkstra"
	v1 "jinli.io/shortestpath/pkg/apis/dijkstra/v1"
	v2 "jinli.io/shortestpath/pkg/apis/dijkstra/v2"
	"k8s.io/apimachinery/pkg/runtime"
	utilruntime "k8s.io/apimachinery/pkg/util/runtime"
)

// 在APIServer包中被引入
func init() {
	Install(dijkstra.Scheme)
}

// Install registers the API group and adds types to a scheme
// 注册API组和类型到Scheme
func Install(scheme *runtime.Scheme) {
	utilruntime.Must(dijkstra.AddToScheme(scheme))
	utilruntime.Must(v1.AddToScheme(scheme))
	// 注册V2版本
	utilruntime.Must(v2.AddToScheme(scheme))
	// 设置版本优先级V2高于V1
	utilruntime.Must(scheme.SetVersionPriority(v2.SchemeGroupVersion, v1.SchemeGroupVersion))
}
