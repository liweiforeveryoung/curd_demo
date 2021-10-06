package cmd

import (
	"curd_demo/api"
	"github.com/spf13/cobra"
)

var httpCommand = cobra.Command{
	Use:   "http",
	Short: "启动 http 服务",
	RunE: func(cmd *cobra.Command, args []string) error {
		return api.StartHttpService()
	},
}
