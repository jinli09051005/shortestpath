package apiserver

import (
	"fmt"

	"jinli.io/shortestpath/pkg/apis/dijkstra"
	"jinli.io/shortestpath/pkg/storage"

	// 引入Install包的初始化函数，注册API group和types到全局scheme
	_ "jinli.io/shortestpath/pkg/apis/dijkstra/install"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apiserver/pkg/registry/generic"
	genericapiserver "k8s.io/apiserver/pkg/server"
)

var (
	// 这些类型是跨版本的，不依赖于特定的API版本，这些类型被称为未版本化类型。
	unversionedVersion = schema.GroupVersion{Group: "", Version: "v1"}
	unversionedTypes   = []runtime.Object{
		&metav1.Status{},
		&metav1.APIVersions{},
		&metav1.APIGroupList{},
		&metav1.APIGroup{},
		&metav1.APIResourceList{},
	}
)

func init() {
	// we need to add the options to empty v1
	// 添加跨版本资源到Scheme
	metav1.AddToGroupVersion(dijkstra.Scheme, unversionedVersion)
	dijkstra.Scheme.AddUnversionedTypes(unversionedVersion, unversionedTypes...)
}

// 安装REST API
func (cs *CustomExtenServer) InstallAPIS(extraConfig *ExtraConfig, optsGetter generic.RESTOptionsGetter, restStorageProviders ...storage.RESTStorageProvider) {
	var apiGroupsInfo []genericapiserver.APIGroupInfo

	for _, restStoragerBuilder := range restStorageProviders {
		groupName := restStoragerBuilder.GroupName()
		apiGroupInfo, enabled := restStoragerBuilder.NewRESTStorage(optsGetter)
		if !enabled {
			fmt.Printf("initializing api group %q error,skiping.", groupName)
			continue
		}
		apiGroupsInfo = append(apiGroupsInfo, apiGroupInfo)
	}

	for index := range apiGroupsInfo {
		// Kubernetes层面安装REST API
		if err := cs.GenericAPIServer.InstallAPIGroup(&apiGroupsInfo[index]); err != nil {
			fmt.Printf("registering api group error: %v.failed.", err)
			panic(err)
		}
	}
}
