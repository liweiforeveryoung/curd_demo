package cmd

import (
	"context"
	"github.com/liweiforeveryoung/curd_demo/dep"
	"time"

	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var migrateCommand = cobra.Command{
	Use:   "migrate",
	Short: "迁移 sql 文件到数据库",
	RunE: func(cmd *cobra.Command, args []string) error {
		logrus.Info("start migrate ...")
		ctx, cancel := context.WithTimeout(context.Background(), time.Minute*2)
		defer cancel()
		err := dep.Hub.DB.Migrate(ctx)
		if err != nil {
			return errors.WithStack(err)
		}
		logrus.Info("end migrate")
		return nil
	},
}
