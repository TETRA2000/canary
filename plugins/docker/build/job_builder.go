package build

import (
	"bytes"
	"context"
	dockertypes "github.com/docker/docker/api/types"
	"github.com/tetra2000/canary/api/job"
	"github.com/tetra2000/canary/api/types"
	pluginTypes "github.com/tetra2000/canary/plugins/docker/api/types"
)

type JobBuilder struct {
	Client *pluginTypes.IDockerClient
}

func (jb *JobBuilder) BuildJob(job job.Job) (types.JobResult, error) {
	// TODO: tags, buildArgs, etc...
	options := dockertypes.ImageBuildOptions{}
	res, err := (*jb.Client).ImageBuild(context.Background(), job.BuildContext, options)
	if err != nil {
		return types.JobResult{}, err
	}

	buf := new(bytes.Buffer)
	_, err = buf.ReadFrom(res.Body)
	if err != nil {
		return types.JobResult{}, err
	}

	// TODO Beautify output
	output := buf.String()

	return types.JobResult{ConsoleOutput: output}, nil
}
