package docker

import (
	"context"
	"fmt"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/network"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use: "docker",
	Short: "Docker control",
	Run: func(cmd *cobra.Command, args []string) {
		Execute()
	},
}

var name string

func init() {
	Cmd.Flags().StringVarP(&name, "name", "n", "", "Container Name")
	Cmd.MarkFlagRequired("name")
}

func Execute() {
	runDocker(name)
}
func runDocker(name string) {
	cli, err := client.NewClientWithOpts(client.FromEnv)
	if err != nil {
		panic(err)
	}

	containers, err := cli.ContainerList(context.Background(), types.ContainerListOptions{})
	if err != nil {
		panic(err)
	}

	for _, container := range containers {
		fmt.Printf("%s %s\n", container.ID[:4], container.Image)
	}
	body, err := cli.ContainerCreate(context.Background(), &container.Config{Image: "python"}, &container.HostConfig{}, &network.NetworkingConfig{}, nil, name)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(body)
}
