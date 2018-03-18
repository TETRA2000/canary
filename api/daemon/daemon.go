package daemon

import (
	"github.com/tetra2000/canary/api/types"
)

type Daemon struct {
	TaskHandlers map[string][]*types.Plugin
}

func (d *Daemon) InvokeTask(taskName string, param types.PluginParam) {
	handlers := d.TaskHandlers[taskName]
	for _, h := range handlers {
		(*h).Exec(taskName, param)
	}
}

func (d *Daemon) GetHandlers(taskName string) []*types.Plugin {
	handlers := d.TaskHandlers[taskName]
	if handlers == nil {
		return []*types.Plugin{}
	} else {
		return handlers
	}
}

func (d *Daemon) RegisterTaskHandler(taskName string, handler *types.Plugin) {
	handlers := d.GetHandlers(taskName)
	if !d.IsHandlerRegistered(taskName, handler) {
		handlers = append(handlers, handler)
		d.TaskHandlers[taskName] = handlers
	}
}

func (d *Daemon) IsHandlerRegistered(taskName string, handler *types.Plugin) bool {
	handlers := d.GetHandlers(taskName)
	for _, h := range handlers {
		if h == handler {
			return true
		}
	}
	return false
}
