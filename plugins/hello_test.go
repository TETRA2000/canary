package main

import (
	"testing"
	"plugin"
	"github.com/tetra2000/canary/api/types"
)

var (
	pluginFile = "./hello.so"
)

func TestHelloPlugin_Exec(t *testing.T) {
	plug, err := plugin.Open(pluginFile)
	if err != nil {
		t.Error(err)
	}

	symPlugin, err := plug.Lookup("Plugin")
	if err != nil {
		t.Error(err)
	}

	var plg types.Plugin
	plg, ok := symPlugin.(types.Plugin)
	if !ok {
		t.Error("unexpected type from module symbol")
	}

	plg.Exec("", types.PluginParam{})
}
