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
	dockerPluginDemo();
}


func pluginDemo() {
	plug, err := plugin.Open("plugins/hello.so")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	symPlugin, err := plug.Lookup("Plugin")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	var plg api.Plugin
	plg, ok := symPlugin.(api.Plugin)
	if !ok {
		fmt.Println("unexpected type from module symbol")
		os.Exit(1)
	}

	plg.Exec("", api.PluginArg{})
}

func dockerPluginDemo() {
	plug, err := plugin.Open("plugins/docker.so")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	symPlugin, err := plug.Lookup("Plugin")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	var plg api.Plugin
	plg, ok := symPlugin.(api.Plugin)
	if !ok {
		fmt.Println("unexpected type from module symbol")
		os.Exit(1)
	}

	plg.Exec("", api.PluginArg{})
}