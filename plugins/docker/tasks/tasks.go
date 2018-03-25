package tasks

import (
	"github.com/tetra2000/canary/api/types"
	dockertypes "github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"fmt"
	"context"
	"github.com/docker/docker/api/types/network"
	"github.com/docker/docker/client"
	//"github.com/docker/docker/pkg/streamformatter"
	"github.com/tetra2000/canary/plugins/docker/build"
	"bytes"
)

// Demo
func ListContainers(param types.PluginParam) types.PluginResult {
	fmt.Print("Listing Docker containers.\n")

	cli, err := client.NewEnvClient()
	if err != nil {
		panic(err)
	}

	containers, err := cli.ContainerList(context.Background(), dockertypes.ContainerListOptions{})
	if err != nil {
		panic(err)
	}

	for _, c := range containers {
		fmt.Printf("%s %s\n", c.ID[:10], c.Image)
	}

	return types.PluginResult{Output: "", Err: nil}
}

func Run(param types.PluginParam) types.PluginResult {
	cli, err := client.NewEnvClient()
	if err != nil {
		panic(err)
	}

	config := &container.Config{
		WorkingDir: param.Workdir,
	}
	hostConfig := &container.HostConfig{}
	networkingConfig := &network.NetworkingConfig{}
	containerName := ""

	c, err := cli.ContainerCreate(context.Background(), config, hostConfig, networkingConfig, containerName)
	if err != nil {
		panic(err)
	}

	fmt.Println(c)

	return types.PluginResult{Output: "", Err: nil}
}

func Build(param types.PluginParam) types.PluginResult {
	ctx := context.Background()
	cli, err := client.NewEnvClient()
	if err != nil {
		panic(err)
	}

	pluginTar := &build.Tar{}
	buildContext, err := pluginTar.ArchiveDirectory(param.Workdir, ".dockerignore")
	if err != nil {
		panic(err)
	}

	options := dockertypes.ImageBuildOptions{}

	res, err := cli.ImageBuild(ctx, buildContext, options)
	if err != nil {
		panic(err)
	}

	// TODO remove
	buf := new(bytes.Buffer)
	buf.ReadFrom(res.Body)
	println(buf.String())

	return types.PluginResult{Output: "", Err: nil}
}
