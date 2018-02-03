package main

import (
	"fmt"
	"os"
	"plugin"
	"github.com/tetra2000/canary/api"
)

func main () {
	fmt.Println("Hello!!")
	pluginDemo()
	dockerPluginDemo()
	gitPluginDemo()
}


func pluginDemo() {
	plg, err := loadPlugin("plugins/hello.so")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	plg.Exec("", api.PluginArg{})
}

func dockerPluginDemo() {
	plg, err := loadPlugin("plugins/docker.so")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	plg.Exec("", api.PluginArg{})
}

func gitPluginDemo() {
	plg, err := loadPlugin("plugins/git.so")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	plg.Exec("", api.PluginArg{})
}

func loadPlugin(path string) (api.Plugin, error) {
	plug, err := plugin.Open(path)
	if err != nil {
		return nil, err
	}

	symPlugin, err := plug.Lookup("Plugin")
	if err != nil {
		return nil, err
	}

	var plg api.Plugin
	plg, ok := symPlugin.(api.Plugin)
	if !ok {
		return nil, err
	}

	return plg, nil
}