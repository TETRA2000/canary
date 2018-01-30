package main

import (
	"fmt"
	"github.com/tetra2000/canary/api"
)

// To suppress warning `relocation target main.main not defined`
func main () {}

type HelloPlugin struct {

}

func (p HelloPlugin) Name() string {
	return "HelloPlugin"
}

func (p HelloPlugin) Exec(taskName string, args api.PluginArg) api.PluginResult {
	fmt.Println("Hello from plugin!!")
	return api.PluginResult{Output: "", Err: nil}
}

var Plugin HelloPlugin
