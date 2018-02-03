package build

import (
	"github.com/docker/docker/builder/dockerignore"
	"os"
)

func ReadIgnoreFile(path string) ([]string, error) {
	ignoreFile, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	return dockerignore.ReadAll(ignoreFile)
}
