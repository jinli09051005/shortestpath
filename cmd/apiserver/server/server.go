/*
server是apiserver的实例
*/
package server

import (
	"jinli.io/shortestpath/cmd/apiserver/config"
	"jinli.io/shortestpath/cmd/apiserver/options"
	customextenserver "jinli.io/shortestpath/pkg/apiserver"
	genericapiserver "k8s.io/apiserver/pkg/server"
)

type APIServer struct {
	serverName string
	server     *genericapiserver.GenericAPIServer
}

func CreateAPIServer(serverName string) *APIServer {
	return &APIServer{
		serverName: serverName,
	}
}

// 创建自定义扩展API Server实例
func (ces *APIServer) CreateServerChain(opts *options.Options) error {
	// 根据options创建普通配置
	cfg, err := config.CreateConfigFromOptions(ces.serverName, opts)
	if err != nil {
		return err
	}
	// 根据普通配置创建RecommendedConfig
	customExtenServerReConfig := createAPIServerReConfig(cfg)
	// 根据RecommendedConfig创建CompletedConfig
	customExtenServerCmConfig := customExtenServerReConfig.Complete()
	// 根据CompletedConfig创建自定义Server
	newCustomExtenServer, err := customExtenServerCmConfig.New()
	if err != nil {
		return err
	}

	newCustomExtenServer.GenericAPIServer.AddPostStartHook("start-dijkstra-api-server-informers", func(context genericapiserver.PostStartHookContext) error {
		customExtenServerCmConfig.GenericConfig.SharedInformerFactory.Start(context.StopCh)
		opts.SharedInformerFactory.Start(context.StopCh)
		return nil
	})

	ces.server = newCustomExtenServer.GenericAPIServer
	return nil
}

// 运行自定义扩展API Server实例
func (ces *APIServer) Run(stopCh <-chan struct{}) error {
	return ces.server.PrepareRun().Run(stopCh)
}

// 根据普通配置创建RecommendedConfig
func createAPIServerReConfig(cfg *config.Config) *customextenserver.RecommendConfig {
	return &customextenserver.RecommendConfig{
		GenericConfig: cfg.GenericAPIServerConfig,
		ExtraConfig: customextenserver.ExtraConfig{
			ServerName: cfg.ServerName,
		},
	}
}
