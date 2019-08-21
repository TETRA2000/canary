package main

import (
	"fmt"
	"os"
	"plugin"
	daemonApi "github.com/tetra2000/canary/api/daemon"
	"github.com/tetra2000/canary/api/types"
)

var daemon *daemonApi.Daemon

func main() {
	daemon = &daemonApi.Daemon{TaskHandlers: make(map[string][]*types.Plugin)}
	loadDefaultPlugins(daemon)

	demo()
}

// canary.demo
// TODO: remove (only for demo)
func demo() {
	var results []types.TaskResult
	results = append(results, daemon.InvokeTask("git:fetch", types.PluginParam{
		Workdir: "./data/projects/demo.canary/repo",
	})...)
	results = append(results, daemon.InvokeTask("docker:build", types.PluginParam{
		Workdir: "./data/projects/demo.canary/repo",
	})...)
	for _, r := range results {
		fmt.Printf("----------\n\nTaskHandler: %s\n\n<Result>\nOutput:\n%s\n\nErr\n%s\n\n", (*r.TaskHandler).Name(), r.Result.Output, r.Result.Err)
	}
}

func loadDefaultPlugins(daemon *daemonApi.Daemon) {
	paths := []string{
		"plugins/docker.so",
		"plugins/git.so"}
	for _, path := range paths {
		plg, err := loadPlugin(path)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		for _, name := range plg.TaskNames() {
			daemon.RegisterTaskHandler(name, &plg)
		}
	}
}

func loadPlugin(path string) (types.Plugin, error) {
	plug, err := plugin.Open(path)
	if err != nil {
		return nil, err
	}

	symPlugin, err := plug.Lookup("Plugin")
	if err != nil {
		return nil, err
	}

	var plg types.Plugin
	plg, ok := symPlugin.(types.Plugin)
	if !ok {
		return nil, err
	}

	return plg, nil
}
