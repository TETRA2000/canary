package main

import (
	"github.com/tetra2000/canary/api/types"
	"github.com/tetra2000/canary/plugins/docker/lib"
	"github.com/tetra2000/canary/plugins/docker/tasks"
	"errors"
)

// To suppress warning `relocation target main.main not defined`
func main () {}

//TODO remove
const version = lib.VERSION

const buildTask = "docker:build"
const runTask = "docker:run"
const previewTask = "docker:preview"

type DockerPlugin struct {

}

func (p DockerPlugin) Name() string {
	return "DockerPlugin"
}

func (p DockerPlugin) TaskNames() []string  {
	return []string{buildTask, runTask, previewTask}
}

func (p DockerPlugin) Exec(taskName string, param types.PluginParam) types.PluginResult {
	switch taskName {
	case buildTask:
		return tasks.Build(param)
		break
	case runTask:
		return tasks.Run(param)
		break
	default:
		return types.PluginResult{Err: errors.New("Undefined task.")}
	}

	return types.PluginResult{Err: errors.New("Unknown error.")}
}

var Plugin DockerPlugin
