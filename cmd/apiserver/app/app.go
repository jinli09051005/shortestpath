package app

import (
	"github.com/spf13/cobra"
	"jinli.io/shortestpath/cmd/apiserver/options"
	"jinli.io/shortestpath/cmd/apiserver/server"
)

// 设置自定义扩展API Server cobra命令
func CustomExtenAPIServer(o *options.Options, stopCh <-chan struct{}) *cobra.Command {
	cmd := &cobra.Command{
		Short: "Launch a custom API Server",
		Long:  "Launch Nexus API Server",
		RunE: func(cmd *cobra.Command, args []string) error {
			if err := o.Complete(); err != nil {
				return err
			}
			if err := o.Validate(); err != nil {
				return err
			}
			apiserver := server.CreateCustemExtenSetrver("jinli-dijkstra-api")
			if err := apiserver.CreateServerChain(o); err != nil {
				return err
			}
			if err := apiserver.Run(stopCh); err != nil {
				return err
			}
			return nil
		},
	}

	// 解析命令行参数
	flags := cmd.Flags()
	// 传递给推荐选项
	o.RecommendedOptions.AddFlags(flags)

	return cmd
}
