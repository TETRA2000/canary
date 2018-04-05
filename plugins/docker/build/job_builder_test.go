package build

import (
	"testing"
	"github.com/tetra2000/canary/api/job"
	"github.com/stretchr/testify/assert"
	uuid2 "github.com/google/uuid"
)

func TestToBuildFromJob(t *testing.T) {
	jobBuilder := JobBuilder{}

	pluginTar := &Tar{}
	buildCtx, err := pluginTar.ArchiveDirectory("./testdata/workdir/", ".dockerignore")
	if err != nil {
		t.Error(err)
	}

	testJob := job.Job{
		Name: "Test",
		Uuid: uuid2.New().String(),
		BuildContext: buildCtx,
	}

	result := jobBuilder.buildJob(testJob)
	assert.NotNil(t, result)
	assert.NotEqual(t, "", result.ConsoleOutput)
	assert.Nil(t, result.Error)
}
