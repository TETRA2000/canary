package api

type PluginArg struct {
	Argc []int
	Argv []string
}

type PluginResult struct {
	Output string
	Err error
}

type Plugin interface {
	Name() string
	Exec(taskName string, args PluginArg) PluginResult
}
