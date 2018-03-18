package types

type PluginParam struct {
	Args    map[string]string
	Workdir string
}

type PluginResult struct {
	Output string
	Err error
}

type Plugin interface {
	Name() string
	TaskNames() []string
	Exec(taskName string, param PluginParam) PluginResult
}
