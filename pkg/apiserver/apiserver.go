package apiserver

import (
	"jinli.io/shortestpath/pkg/registry/dijkstra/rest"
	"jinli.io/shortestpath/pkg/storage"
	genericapiserver "k8s.io/apiserver/pkg/server"
)

type CustomExtenServer struct {
	GenericAPIServer *genericapiserver.GenericAPIServer
}

func (c CompletedConfig) New() (*CustomExtenServer, error) {
	// 完整配置生成无代理自定义服务,即自定义服务不代理别的服务
	s, err := c.GenericConfig.New(c.ExtraConfig.ServerName, genericapiserver.NewEmptyDelegate())
	if err != nil {
		return nil, err
	}

	cs := &CustomExtenServer{
		GenericAPIServer: s,
	}
	restStorageProviders := []storage.RESTStorageProvider{
		&rest.StorageProvider{
			ClientConfig: c.GenericConfig.LoopbackClientConfig,
		},
	}
	// 安装扩展REST API服务
	cs.InstallAPIS(c.GenericConfig.RESTOptionsGetter, restStorageProviders...)
	return cs, nil
}
