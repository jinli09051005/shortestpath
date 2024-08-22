package rest

import (
	dijkstraclient "jinli.io/shortestpath/generated/client/clientset/internalversion/typed/dijkstra/internalversion"
	"jinli.io/shortestpath/pkg/apis/dijkstra"
	dijkstrav1 "jinli.io/shortestpath/pkg/apis/dijkstra/v1"
	dijkstrav2 "jinli.io/shortestpath/pkg/apis/dijkstra/v2"
	dpstorage "jinli.io/shortestpath/pkg/registry/dijkstra/display"
	knstorage "jinli.io/shortestpath/pkg/registry/dijkstra/knownnodes"
	"jinli.io/shortestpath/pkg/storage"
	"k8s.io/apiserver/pkg/registry/generic"
	"k8s.io/apiserver/pkg/registry/rest"
	genericapiserver "k8s.io/apiserver/pkg/server"
	restclient "k8s.io/client-go/rest"
)

type StorageProvider struct {
	ClientConfig *restclient.Config
}

var _ storage.RESTStorageProvider = &StorageProvider{}

func (s *StorageProvider) GroupName() string {
	return dijkstra.Group
}

func (s *StorageProvider) NewRESTStorage(optsGetter generic.RESTOptionsGetter) (genericapiserver.APIGroupInfo, bool) {
	apiGroupInfo := genericapiserver.NewDefaultAPIGroupInfo(dijkstra.Group, dijkstra.Scheme, dijkstra.ParameterCodec, dijkstra.Codecs)
	apiGroupInfo.VersionedResourcesStorageMap[dijkstrav1.SchemeGroupVersion.Version] = s.v1Storage(optsGetter, s.ClientConfig)
	apiGroupInfo.VersionedResourcesStorageMap[dijkstrav2.SchemeGroupVersion.Version] = s.v2Storage(optsGetter, s.ClientConfig)

	return apiGroupInfo, true
}

func (s *StorageProvider) v1Storage(optsGetter generic.RESTOptionsGetter, clientConfig *restclient.Config) map[string]rest.Storage {
	dijkstraClient := dijkstraclient.NewForConfigOrDie(clientConfig)

	storageMap := map[string]rest.Storage{}
	knREST := knstorage.RESTInPeace(knstorage.NewStorage(dijkstra.Scheme, optsGetter, dijkstraClient))
	dpREST := dpstorage.RESTInPeace(dpstorage.NewStorage(dijkstra.Scheme, optsGetter, dijkstraClient))

	storageMap["knownnodeses"] = knREST.KnownNodes
	// 添加逻辑子资源，即更新子资源不影响总体
	storageMap["knownnodeses/status"] = knREST.Status
	storageMap["displays"] = dpREST.Display
	// 添加逻辑子资源，即更新子资源不影响总体
	storageMap["displays/status"] = dpREST.Status
	return storageMap
}

func (s *StorageProvider) v2Storage(optsGetter generic.RESTOptionsGetter, clientConfig *restclient.Config) map[string]rest.Storage {
	dijkstraClient := dijkstraclient.NewForConfigOrDie(clientConfig)

	storageMap := map[string]rest.Storage{}
	knREST := knstorage.RESTInPeace(knstorage.NewStorage(dijkstra.Scheme, optsGetter, dijkstraClient))
	dpREST := dpstorage.RESTInPeace(dpstorage.NewStorage(dijkstra.Scheme, optsGetter, dijkstraClient))

	storageMap["knownnodeses"] = knREST.KnownNodes
	// 添加逻辑子资源，即更新子资源不影响总体
	storageMap["knownnodeses/status"] = knREST.Status
	storageMap["displays"] = dpREST.Display
	// 添加逻辑子资源，即更新子资源不影响总体
	storageMap["displays/status"] = dpREST.Status
	return storageMap
}
