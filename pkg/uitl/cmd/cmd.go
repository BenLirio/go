package cmd

import (

	"github.com/BenLirio/op/pkg/util/download"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	rootCmd.AddCommand(download.Cmd)
}

var rootCmd = &cobra.Command{
	Use: "util",
	Short: "Personal utilities",
}
