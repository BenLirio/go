package watch

import (
	"os"
	"fmt"
	"time"
	"os/exec"
	"strings"

	"github.com/BenLirio/op/pkg/ai/watch/walk"

	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use: "watch",
	Short: "watch files",
	Run: func(cmd *cobra.Command, args []string) {
		Execute()
	},
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

var folder string
var script string

func init() {
	Cmd.Flags().StringVarP(&folder, "folder", "f", ".", "Folder to watch")
	Cmd.Flags().StringVarP(&script, "script", "s", "", "Script to run")
	Cmd.MarkFlagRequired("script")
}

func clear() error {
	goExecPath, err := exec.LookPath("clear")
	if err != nil {
		return err
	}
	onChangeScript := &exec.Cmd {
		Path: goExecPath,
		Args: []string{"clear"},
		Stdout: os.Stdout,
		Stderr: os.Stderr,
	}
	onChangeScript.Run()
	return nil
}

func runScript() error {
	args := strings.Split(script, " ")
	goExecPath, err := exec.LookPath(args[0])
	if err != nil {
		return err
	}
	onChangeScript := &exec.Cmd {
		Path: goExecPath,
		Args: args,
		Stdout: os.Stdout,
		Stderr: os.Stderr,
	}
	onChangeScript.Run()
	return nil
}


func Execute() {
	err := walk.GetFileStats(folder)
	check(err)
	err = clear()
	check(err)
	now := time.Now()
	fmt.Print("[", now.Format("07:04:05"), "]")
	fmt.Println(" Watching:", folder)
	for {
		diff, err := walk.CheckDiff()
		check(err)
		if diff {
			err := clear()
			check(err)
			now := time.Now()
			fmt.Println("[", now.Format("07:04:05"), "]")
			err = runScript()
			check(err)
		}
		time.Sleep(500 * time.Millisecond)
	}
}
