package cmd

import (

	"github.com/BenLirio/op/pkg/ai/version"
	"github.com/BenLirio/op/pkg/ai/yaml"
	"github.com/BenLirio/op/pkg/ai/browse"
	"github.com/BenLirio/op/pkg/ai/watch"
	"github.com/BenLirio/op/pkg/ai/docker"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	viper.SetDefault("host", "localhost")
	viper.SetDefault("port", "8080")
	rootCmd.AddCommand(version.Cmd)
	rootCmd.AddCommand(yaml.Cmd)
	rootCmd.AddCommand(browse.Cmd)
	rootCmd.AddCommand(watch.Cmd)
	rootCmd.AddCommand(docker.Cmd)
}

var rootCmd = &cobra.Command{
	Use: "ai",
	Short: "Personal tool to increase productivity",
}
