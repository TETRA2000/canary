package main

import (
	"fmt"
	"github.com/tetra2000/canary/api"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
	"context"
)

// To suppress warning `relocation target main.main not defined`
func main () {}

type DockerPlugin struct {

}

func (p DockerPlugin) Name() string {
	return "DockerPlugin"
}

func (p DockerPlugin) Exec(arg api.PluginArg) api.PluginResult {
	fmt.Print("Listing Docker images.\n")

	cli, err := client.NewEnvClient()
	if err != nil {
		panic(err)
	}

	containers, err := cli.ContainerList(context.Background(), types.ContainerListOptions{})
	if err != nil {
		panic(err)
	}

	for _, container := range containers {
		fmt.Printf("%s %s\n", container.ID[:10], container.Image)
	}

	return api.PluginResult{Output: "", Err: nil}
}

var Plugin DockerPlugin
