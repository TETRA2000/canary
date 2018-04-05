package job

import (
	"io"
)

type Job struct {
	Name string // Human readable name
	Uuid string // UUID as a unique identifier
	BuildContext io.ReadCloser // ReadCloser includes a build context for docker
	// TODO add information of scheduled tasks
}
