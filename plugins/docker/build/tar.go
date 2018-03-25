package build

import (
	"github.com/docker/docker/pkg/archive"
	"io"
	"path/filepath"
)

type Tar struct {

}

func (t *Tar) ArchiveDirectory(workdir string, ignoreFile string) (io.ReadCloser, error) {
	ignoreFilePath := filepath.Join(workdir, ignoreFile)

	ignore := &Ignore{}
	excludes, err := ignore.ReadIgnoreFile(ignoreFilePath)
	if err != nil {
		return nil, err
	}

	tarOptions := &archive.TarOptions{
		ExcludePatterns: excludes,
	}

	return archive.TarWithOptions(workdir, tarOptions)
}
