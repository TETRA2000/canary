package tasks

import (
	"github.com/tetra2000/canary/api/types"
	"gopkg.in/src-d/go-git.v4"
)

func Fetch(param types.PluginParam) types.PluginResult {
	r, err := git.PlainOpen(param.Workdir)
	if err != nil {
		return types.PluginResult{Output: "", Err: err}
	}

	opt := &git.FetchOptions{
		RemoteName: "",
		RefSpecs:   nil,
		Depth:      0,
		Auth:       nil,
		Progress:   nil,
		Tags:       0,
		Force:      false,
	}
	err = r.Fetch(opt)
	if err != nil {
		return types.PluginResult{Output: "", Err: err}
	}

	return types.PluginResult{Output: "", Err: nil}
}

// Ref: https://github.com/src-d/go-git/blob/master/_examples/log/main.go
