package cmd

import (
	"github.com/spf13/cobra"

	"github.com/liweiforeveryoung/curd_demo/api"
)

var httpCommand = cobra.Command{
	Use:   "http",
	Short: "启动 http 服务",
	RunE: func(cmd *cobra.Command, args []string) error {
		return api.StartHttpService()
	},
}
