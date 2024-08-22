package v2

import (
	"k8s.io/apimachinery/pkg/runtime"
)

// 在register中被注册
func addDefaultingFuncs(scheme *runtime.Scheme) error {
	// zz_generated.defaults.go中定义的函数
	return RegisterDefaults(scheme)
}

// 设置默认值函数，在生成的zz_generated.defaults.go中被调用
// 设置什么对象传什么对象，比如这里设置DisplaySpec，参数便是DisplaySpec类型
func SetDefaults_KnownNodesSpec(obj *KnownNodesSpec) {
	if obj.NodeIdentity == "" {
		obj.NodeIdentity = "jinli-default-nodeid"
	}
}

func SetDefaults_DisplaySpec(obj *DisplaySpec) {
	if obj.NodeIdentity == "" {
		obj.NodeIdentity = "jinli-default-nodeid"
	}

	// 默认dijkstra算法
	if obj.Algorithm == "" {
		obj.Algorithm = "dijkstra"
	}
}

func SetDefaults_DisplayStatus(obj *DisplayStatus) {
	// 默认Wait
	if obj.ComputeStatus == "" {
		obj.ComputeStatus = "Wait"
	}
}
