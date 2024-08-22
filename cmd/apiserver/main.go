package main

import (
	"flag"

	"jinli.io/shortestpath/cmd/apiserver/app"
	"jinli.io/shortestpath/cmd/apiserver/options"
	genericapiserver "k8s.io/apiserver/pkg/server"
	"k8s.io/component-base/logs"
	"k8s.io/klog"
)

func main() {
	logs.InitLogs()
	defer logs.FlushLogs()

	// 创建停止通道
	stopCh := genericapiserver.SetupSignalHandler()
	// 创建自定义扩展API Server选项
	opts := options.CreateOptions()
	// 创建自定义扩展API Server启动命令
	cmd := app.CustomExtenAPIServer(opts, stopCh)
	cmd.Flags().AddGoFlagSet(flag.CommandLine)
	// 最终启动自定义扩展API Server
	if err := cmd.Execute(); err != nil {
		klog.Fatal(err)
	}
}
