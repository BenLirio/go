package version

import (
	"fmt"

	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use: "version",
	Short: "Get Version",
	Run: func(cmd *cobra.Command, args []string) {
		Execute()
	},
}

func Execute() {
	fmt.Println("v0.0.0")
}
