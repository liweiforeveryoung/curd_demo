package cmd

import (
	"github.com/liweiforeveryoung/curd_demo/config"
	"github.com/liweiforeveryoung/curd_demo/dep"
	"math/rand"
	"time"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "curd_demo",
	Short: "a curd demo to learn how to write testable code",
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	cobra.OnInitialize(initialize)
	rootCmd.AddCommand(&migrateCommand)
	rootCmd.AddCommand(&httpCommand)
	cobra.CheckErr(rootCmd.Execute())
}

func initialize() {
	config.Initialize()
	dep.Prepare()
	rand.Seed(time.Now().Unix())
}
