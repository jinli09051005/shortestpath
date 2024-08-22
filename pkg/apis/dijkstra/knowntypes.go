package dijkstra

import (
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

const Group = "dijkstra.jinli.io"

var SchemeGroupVersion = schema.GroupVersion{
	Group:   Group,
	Version: runtime.APIVersionInternal,
}

func addKnownTypes(scheme *runtime.Scheme) error {
	// 内部资源类型注册到scheme
	scheme.AddKnownTypes(SchemeGroupVersion,
		&KnownNodes{},
		&KnownNodesList{},
		&Display{},
		&DisplayList{},
	)
	return nil
}

// Kind takes an unqualified kind and returns back a IdentityProvider qualified GroupKind
func Kind(kind string) schema.GroupKind {
	return SchemeGroupVersion.WithKind(kind).GroupKind()
}
func Resource(resource string) schema.GroupResource {
	return SchemeGroupVersion.WithResource(resource).GroupResource()
}
