package build

import (
	"github.com/docker/docker/pkg/archive"
	"io"
	"path/filepath"
)

func ArchiveDirectory(workdir string, ignoreFile string) (io.ReadCloser, error) {
	ignoreFilePath := filepath.Join(workdir, ignoreFile)

	excludes, err := ReadIgnoreFile(ignoreFilePath)
	if err != nil {
		return nil, err
	}

	tarOptions := &archive.TarOptions{
		ExcludePatterns: excludes,
	}

	return archive.TarWithOptions(workdir, tarOptions)
}
