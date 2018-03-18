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
	//ctx := context.Background()
	//cli, err := client.NewEnvClient()
	//if err != nil {
	//	panic(err)
	//}


	// cli.ImageBuild(ctx)


	return types.PluginResult{Output: "", Err: nil}
}

//func Build(param types.PluginParam) types.PluginResult {
//  var (
//    buildCtx      io.ReadCloser
//    //dockerfileCtx io.ReadCloser
//    err           error
//    contextDir    string
//    //tempDir       string
//    relDockerfile string
//    progBuff      io.Writer
//    //buildBuff     io.Writer
//    //remote        string
//  )
//
//  cli, err := client.NewEnvClient()
//  if err != nil {
//    panic(err)
//  }
//
//  contextDir, relDockerfile, err = clibuild.GetContextFromLocalDir(param.Workdir, "Dockerfile")
//
//  // read from a directory into tar archive
//  excludes, err := clibuild.ReadDockerignore(contextDir)
//  if err != nil {
//    return types.PluginResult{Err: err}
//  }
//
//  if err := clibuild.ValidateContextDirectory(contextDir, excludes); err != nil {
//    return types.PluginResult{Err: errors.New(fmt.Sprint("error checking context: '%s'.", err))}
//  }
//
//  excludes = clibuild.TrimBuildFilesFromExcludes(excludes, relDockerfile, false)
//  buildCtx, err = archive.TarWithOptions(contextDir, &archive.TarOptions{
//    ExcludePatterns: excludes,
//    ChownOpts:       &idtools.IDPair{UID: 0, GID: 0},
//  })
//  if err != nil {
//    return types.PluginResult{Err: err}
//  }
//
//  // Setup an upload progress bar
//  progBuff = bytes.NewBuffer(nil)
//  progressOutput := streamformatter.NewProgressOutput(progBuff)
//
//  var body io.Reader
//  if buildCtx != nil {
//    body = progress.NewProgressReader(buildCtx, progressOutput, 0, "", "Sending build context to Docker daemon")
//  }
//
//  buildOptions := dockertypes.ImageBuildOptions{}
//
//  response, err := cli.ImageBuild(context.Background(), body, buildOptions)
//  if err != nil {
//    return types.PluginResult{Err: err}
//  }
//  defer response.Body.Close()
//
//  return types.PluginResult{Output: "", Err: nil}
//}
