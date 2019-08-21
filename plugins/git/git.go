package main

import (
	"errors"
	"github.com/tetra2000/canary/api/types"
	"github.com/tetra2000/canary/plugins/git/lib"
	"github.com/tetra2000/canary/plugins/git/tasks"
)

// To suppress warning `relocation target main.main not defined`
func main () {}

//TODO remove
const version = lib.VERSION
const fetchTask = "git:fetch"

type GitPlugin struct {

}

func (p GitPlugin) Name() string {
	return "GitPlugin"
}

func (p GitPlugin) TaskNames() []string  {
	return []string{fetchTask}
}

func (p GitPlugin) Exec(taskName string, param types.PluginParam) types.PluginResult {
	switch taskName {
	case fetchTask:
		return tasks.Fetch(param)
	default:
		return types.PluginResult{Err: errors.New("undefined task")}
	}
}

var Plugin GitPlugin
