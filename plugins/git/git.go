package main

import (
	"fmt"
	"github.com/tetra2000/canary/api/types"
	"github.com/tetra2000/canary/plugins/git/lib"
	"gopkg.in/src-d/go-git.v4/plumbing/object"
	"gopkg.in/src-d/go-git.v4"
)

// To suppress warning `relocation target main.main not defined`
func main () {}

//TODO remove
const version = lib.VERSION

type GitPlugin struct {

}

func (p GitPlugin) Name() string {
	return "GitPlugin"
}

func (p GitPlugin) TaskNames() []string  {
	return []string{"git:pull"}
}

func (p GitPlugin) Exec(taskName string, param types.PluginParam) types.PluginResult {
	// from https://github.com/src-d/go-git/blob/master/_examples/log/main.go

	fmt.Println("opening repo in {0}", param.Workdir)
	r, err := git.PlainOpen(param.Workdir)
	if err != nil {
		panic(err)
	}

	// Gets the HEAD history from HEAD, just like does:
	fmt.Println("git log")

	// ... retrieves the branch pointed by HEAD
	ref, err := r.Head()
	if err != nil {
		panic(err)
	}

	// ... retrieves the commit history
	cIter, err := r.Log(&git.LogOptions{From: ref.Hash()})
	if err != nil {
		panic(err)
	}

	// ... just iterates over the commits, printing it
	err = cIter.ForEach(func(c *object.Commit) error {
		fmt.Println(c)

		return nil
	})
	if err != nil {
		panic(err)
	}

	return types.PluginResult{Output: "", Err: nil}
}

var Plugin GitPlugin
