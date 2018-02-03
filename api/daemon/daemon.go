package daemon

import (
	"github.com/tetra2000/canary/api/types"
)

type Daemon struct {
}

func (d *Daemon) InvokeTask(taskName string, args types.PluginArg) {
	panic("Not implemented.")
}

func (d *Daemon) RegisterTaskHandler(taskName string, plugin *types.Plugin) {
	panic("Not implemented.")
}
