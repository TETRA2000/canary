package main

import (
	"fmt"
	"os"
	"plugin"
	"github.com/tetra2000/canary/api/types"
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
	plg.Exec("", types.PluginArg{})
}

func dockerPluginDemo() {
	plg, err := loadPlugin("plugins/docker.so")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	plg.Exec("", types.PluginArg{})
}

func gitPluginDemo() {
	plg, err := loadPlugin("plugins/git.so")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	plg.Exec("", types.PluginArg{})
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