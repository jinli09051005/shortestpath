package openapi

import (
	generatedopenapi "jinli.io/shortestpath/generated/openapi"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/util/sets"
	openapinamer "k8s.io/apiserver/pkg/endpoints/openapi"
	genericapiserver "k8s.io/apiserver/pkg/server"
	"k8s.io/apiserver/pkg/server/filters"
	utilopenapi "k8s.io/apiserver/pkg/util/openapi"
	"k8s.io/component-base/version"
	"k8s.io/kube-openapi/pkg/validation/spec"
)

func SetupOpenAPI(scheme *runtime.Scheme, genericAPIServerConfig *genericapiserver.Config, title, license string) {
	// wrap the definitions to revert any changes from disabled features
	getOpenAPIDefinitions := utilopenapi.GetOpenAPIDefinitionsWithoutDisabledFeatures(generatedopenapi.GetOpenAPIDefinitions)
	namer := openapinamer.NewDefinitionNamer(scheme)
	kubeVersion := version.Get()
	genericAPIServerConfig.Version = &kubeVersion
	// OpenAPI v2.0
	genericAPIServerConfig.OpenAPIConfig = genericapiserver.DefaultOpenAPIConfig(getOpenAPIDefinitions, namer)
	genericAPIServerConfig.OpenAPIConfig.Info.Title = title
	genericAPIServerConfig.OpenAPIConfig.Info.License = &spec.License{Name: license}
	// OpenAPI v3.0
	genericAPIServerConfig.OpenAPIV3Config = genericapiserver.DefaultOpenAPIConfig(getOpenAPIDefinitions, namer)
	genericAPIServerConfig.OpenAPIV3Config.Info.Title = title
	genericAPIServerConfig.OpenAPIV3Config.Info.License = &spec.License{Name: license}
	genericAPIServerConfig.LongRunningFunc = filters.BasicLongRunningRequestCheck(
		sets.NewString("watch", "proxy"),
		sets.NewString("attach", "exec", "proxy", "log", "portforward"),
	)
}
