package tasks

import (
	"github.com/tetra2000/canary/api/types"
	dockertypes "github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"fmt"
	"context"
	"github.com/docker/docker/api/types/network"
	"github.com/tetra2000/canary/plugins/docker/build"
	"bytes"
	"github.com/tetra2000/canary/plugins/docker/client"
	reacClient "github.com/docker/docker/client"
)

// Demo
func ListContainers(param types.PluginParam) types.PluginResult {
	// TODO replace with client.NewDockerClient
	fmt.Print("Listing Docker containers.\n")

	cli, err := reacClient.NewEnvClient()
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

// TODO Replace with build + start
func Run(param types.PluginParam) types.PluginResult {
	// TODO replace with client.NewDockerClient
	cli, err := reacClient.NewEnvClient()
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
	cli, err := client.NewDockerClient()
	if err != nil {
		return types.PluginResult{Output: "", Err: err}
	}

	pluginTar := &build.Tar{}
	buildContext, err := pluginTar.ArchiveDirectory(param.Workdir, ".dockerignore")
	if err != nil {
		return types.PluginResult{Output: "", Err: err}
	}

	options := dockertypes.ImageBuildOptions{}

	res, err := cli.ImageBuild(ctx, buildContext, options)
	if err != nil {
		return types.PluginResult{Output: "", Err: err}
	}

	buf := new(bytes.Buffer)
	buf.ReadFrom(res.Body)
	return types.PluginResult{Output: buf.String(), Err: nil}
}
