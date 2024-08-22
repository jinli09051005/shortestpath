package storage

import (
	"k8s.io/apiserver/pkg/registry/generic"
	genericserver "k8s.io/apiserver/pkg/server"
)

type RESTStorageProvider interface {
	NewRESTStorage(optsGetter generic.RESTOptionsGetter) (genericserver.APIGroupInfo, bool)
	GroupName() string
}
