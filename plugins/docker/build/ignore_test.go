package build

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestToReadDockerignore(t *testing.T) {
	ignore := &Ignore{}
	excludes, err := ignore.ReadIgnoreFile("./testdata/workdir/.dockerignore")
	if err != nil {
		t.Error(err)
	}

	assert.Contains(t, excludes, "log")
	assert.Contains(t, excludes, "exclud.txt")
}
