package config

import (
	"jinli.io/shortestpath/cmd/apiserver/options"
	"jinli.io/shortestpath/pkg/apis/dijkstra"
	"jinli.io/shortestpath/pkg/openapi"
	genericapiserver "k8s.io/apiserver/pkg/server"
)

const (
	title   = "Jinli ShortestPath Dijkstra API"
	license = "Apache 2.0"
)

type Config struct {
	ServerName             string
	GenericAPIServerConfig *genericapiserver.RecommendedConfig
}

func CreateConfigFromOptions(ServerName string, opts *options.Options) (*Config, error) {
	// 创建genericAPIServerConfig
	genericAPIServerConfig := genericapiserver.NewRecommendedConfig(dijkstra.Codecs)
	// opts设置genericAPIServerConfig设置，设置admission等
	if err := opts.RecommendedOptions.ApplyTo(genericAPIServerConfig); err != nil {
		return nil, err
	}
	// 设置genericAPIServerConfig OpenAPI配置
	openapi.SetupOpenAPI(dijkstra.Scheme, &genericAPIServerConfig.Config, title, license)
	// 返回配置
	cfg := &Config{
		ServerName:             ServerName,
		GenericAPIServerConfig: genericAPIServerConfig,
	}

	return cfg, nil
}
