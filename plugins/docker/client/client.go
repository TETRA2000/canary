package client

import (
	"github.com/docker/docker/client"
	"github.com/tetra2000/canary/plugins/docker/api/types"
)

func NewDockerClient() (*types.DockerClient, error) {
	cli, err := client.NewEnvClient()
	if err != nil {
		return nil, err
	}

	return &types.DockerClient{
		RealClient: cli,
	}, nil
}
