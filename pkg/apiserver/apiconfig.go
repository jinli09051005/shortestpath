package apiserver

import (
	"k8s.io/apimachinery/pkg/version"
	genericapiserver "k8s.io/apiserver/pkg/server"
)

// 推荐配置
type RecommendConfig struct {
	// 通用配置
	GenericConfig *genericapiserver.RecommendedConfig
	// 额外配置
	ExtraConfig ExtraConfig
}

// 额外配置
type ExtraConfig struct {
	ServerName string
}

// 私有完整配置
type completedConfig struct {
	// 由推荐配置生成
	GenericConfig genericapiserver.CompletedConfig
	ExtraConfig   *ExtraConfig
}

// 完整配置
type CompletedConfig struct {
	*completedConfig
}

// 推荐配置生成完整配置
func (cfg *RecommendConfig) Complete() CompletedConfig {
	c := completedConfig{
		cfg.GenericConfig.Complete(),
		&cfg.ExtraConfig,
	}

	c.GenericConfig.Version = &version.Info{
		Major: "1",
		Minor: "0",
	}

	return CompletedConfig{&c}
}
