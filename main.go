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

	daemon.InvokeTask("hello", types.PluginParam{})

	demo()
}

// canary.demo
func demo() {
	daemon.InvokeTask("git:pull", types.PluginParam{
		Workdir: "/opt/canary/data/projects/demo.canary/repo",
	})
	daemon.InvokeTask("docker:build", types.PluginParam{
		Workdir: "/opt/canary/data/projects/demo.canary",
	})

	daemon.InvokeTask("docker:build", types.PluginParam{
		Workdir: "/opt/canary/data/projects/demo.canary",
	})
}

func loadDefaultPlugins(daemon *daemonApi.Daemon) {
	paths := []string{"plugins/hello.so",
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
