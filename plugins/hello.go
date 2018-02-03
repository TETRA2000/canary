package main

import (
	"fmt"
	"github.com/tetra2000/canary/api/types"
)

// To suppress warning `relocation target main.main not defined`
func main () {}

type HelloPlugin struct {

}

func (p HelloPlugin) Name() string {
	return "HelloPlugin"
}

func (p HelloPlugin) TaskNames() []string  {
	return []string{"hello"}
}

func (p HelloPlugin) Exec(taskName string, args types.PluginArg) types.PluginResult {
	fmt.Println("Hello from plugin!!")
	return types.PluginResult{Output: "", Err: nil}
}

var Plugin HelloPlugin
