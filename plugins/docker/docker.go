package main

import (
	"fmt"
	"github.com/tetra2000/canary/api/types"
	"github.com/tetra2000/canary/plugins/docker/lib"

	dockerTypes "github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
	"context"
)

// To suppress warning `relocation target main.main not defined`
func main () {}

//TODO remove
const version = lib.VERSION

type DockerPlugin struct {

}

func (p DockerPlugin) Name() string {
	return "DockerPlugin"
}

func (p DockerPlugin) TaskNames() []string  {
	return []string{"docker:build", "docker:run", "docker:preview"}
}

func (p DockerPlugin) Exec(taskName string, args types.PluginArg) types.PluginResult {
	fmt.Print("Listing Docker containers.\n")

	cli, err := client.NewEnvClient()
	if err != nil {
		panic(err)
	}

	containers, err := cli.ContainerList(context.Background(), dockerTypes.ContainerListOptions{})
	if err != nil {
		panic(err)
	}

	for _, container := range containers {
		fmt.Printf("%s %s\n", container.ID[:10], container.Image)
	}

	return types.PluginResult{Output: "", Err: nil}
}

var Plugin DockerPlugin
