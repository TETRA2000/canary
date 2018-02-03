package main

import (
	"fmt"
	"github.com/tetra2000/canary/api/types"
	"github.com/tetra2000/canary/plugins/git/lib"
	"gopkg.in/src-d/go-git.v4/plumbing/object"
	"gopkg.in/src-d/go-git.v4"
	"gopkg.in/src-d/go-git.v4/storage/memory"
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

func (p GitPlugin) Exec(taskName string, args types.PluginArg) types.PluginResult {
	// from https://github.com/src-d/go-git/blob/master/_examples/log/main.go

	fmt.Println("git clone https://github.com/src-d/go-siva")
	r, err := git.Clone(memory.NewStorage(), nil, &git.CloneOptions{
		URL: "https://github.com/src-d/go-siva",
	})
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
