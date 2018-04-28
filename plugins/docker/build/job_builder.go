package build

import (
	"github.com/tetra2000/canary/api/job"
	"github.com/tetra2000/canary/api/types"
	pluginTypes "github.com/tetra2000/canary/plugins/docker/api/types"
)

type JobBuilder struct {
	Client *pluginTypes.IDockerClient
}

func (jb *JobBuilder) buildJob(job job.Job) types.JobResult {
	return types.JobResult{}
}
