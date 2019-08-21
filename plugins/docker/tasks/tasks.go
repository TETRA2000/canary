package tasks

import (
	"bytes"
	"context"
	"fmt"
	dockertypes "github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/network"
	realClient "github.com/docker/docker/client"
	"github.com/google/uuid"
	"github.com/tetra2000/canary/api/job"
	"github.com/tetra2000/canary/api/types"
	pluginTypes "github.com/tetra2000/canary/plugins/docker/api/types"
	"github.com/tetra2000/canary/plugins/docker/build"
	"github.com/tetra2000/canary/plugins/docker/client"
)

// Demo
func ListContainers(param types.PluginParam) types.PluginResult {
	// TODO replace with client.NewDockerClient
	fmt.Print("Listing Docker containers.\n")

	cli, err := realClient.NewEnvClient()
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
	cli, err := realClient.NewEnvClient()
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

func BuildJob(param types.PluginParam) types.PluginResult {
	var cli pluginTypes.IDockerClient
	var err error
	cli, err= client.NewDockerClient()
	if err != nil {
		return types.PluginResult{Output: "", Err: err}
	}

	jobBuilder := build.JobBuilder{
		Client: &cli,
	}

	pluginTar := &build.Tar{}
	// TODO: allow to configure .dockerignore filename.
	buildCtx, err := pluginTar.ArchiveDirectory(param.Workdir, ".dockerignore")
	if err != nil {
		return types.PluginResult{Output: "", Err: err}
	}

	testJob := job.Job{
		Name:         "Test",
		Uuid:         uuid.New().String(),
		BuildContext: buildCtx,
	}

	result, err := jobBuilder.BuildJob(testJob)
	if err != nil {
		return types.PluginResult{Output: "", Err: err}
	}

	return types.PluginResult{Output: result.ConsoleOutput, Err: nil}
}

// Deprecated: Use BuildJob instead.
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
	_, err = buf.ReadFrom(res.Body)
	if err != nil {
		return types.PluginResult{Output: "", Err: err}
	}
	return types.PluginResult{Output: buf.String(), Err: nil}
}
