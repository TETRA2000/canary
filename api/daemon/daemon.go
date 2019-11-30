package daemon

import (
	"fmt"
	"github.com/tetra2000/canary/api/types"
)

type Daemon struct {
	TaskHandlers map[string][]*types.Plugin
}

// Invoke registered tasks for the given name.
func (d *Daemon) InvokeTask(taskName string, param types.PluginParam) []types.TaskResult {
	if !d.IsKnownHandler(taskName) {
		fmt.Printf("Unknown handler %s\n", taskName)
		panic(0)
	}

	var results []types.TaskResult
	handlers := d.TaskHandlers[taskName]
	for _, h := range handlers {
		results = append(results, types.TaskResult{TaskHandler: h, Result: (*h).Exec(taskName, param)})
	}
	return results
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

func (d *Daemon) IsKnownHandler(taskName string) bool {
	handlers := d.GetHandlers(taskName)
	return len(handlers) > 0
}
