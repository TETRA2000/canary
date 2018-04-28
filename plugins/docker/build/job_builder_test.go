package build

import (
	"testing"
	"github.com/tetra2000/canary/api/job"
	"github.com/stretchr/testify/assert"
	uuid2 "github.com/google/uuid"
    "github.com/tetra2000/canary/plugins/docker/api/types"
	"github.com/docker/docker/api/types/image"
	"github.com/docker/docker/api/types/registry"
	"github.com/docker/docker/api/types/filters"
	"golang.org/x/net/context"
	"io"
	dockerTypes "github.com/docker/docker/api/types"
)

func TestToBuildFromJob(t *testing.T) {
	cli, err := newMockClient()
	if err != nil {
		t.Error(err)
	}

	jobBuilder := JobBuilder{
		Client: cli,
	}

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

type MockDockerClient struct {
	types.IDockerClient
}

func newMockClient() (*types.IDockerClient, error) {
	var cli types.IDockerClient = &MockDockerClient{

	}
	return &cli, nil
}

func (cli *MockDockerClient) ImageBuild(ctx context.Context, context io.Reader, options dockerTypes.ImageBuildOptions) (dockerTypes.ImageBuildResponse, error) {
	return dockerTypes.ImageBuildResponse{}, nil
}

func (cli *MockDockerClient) BuildCachePrune(ctx context.Context) (*dockerTypes.BuildCachePruneReport, error) {
	return nil, nil
}

func (cli *MockDockerClient) ImageCreate(ctx context.Context, parentReference string, options dockerTypes.ImageCreateOptions) (io.ReadCloser, error){
	return nil, nil
}

func (cli *MockDockerClient) ImageHistory(ctx context.Context, image string) ([]image.HistoryResponseItem, error) {
	return nil, nil
}

func (cli *MockDockerClient) ImageImport(ctx context.Context, source dockerTypes.ImageImportSource, ref string, options dockerTypes.ImageImportOptions) (io.ReadCloser, error) {
	return nil, nil
}

func (cli *MockDockerClient) ImageInspectWithRaw(ctx context.Context, image string) (dockerTypes.ImageInspect, []byte, error) {
	return dockerTypes.ImageInspect{}, nil, nil
}

func (cli *MockDockerClient) ImageList(ctx context.Context, options dockerTypes.ImageListOptions) ([]dockerTypes.ImageSummary, error) {
	return nil, nil
}

func (cli *MockDockerClient) ImageLoad(ctx context.Context, input io.Reader, quiet bool) (dockerTypes.ImageLoadResponse, error) {
	return dockerTypes.ImageLoadResponse{}, nil
}

func (cli *MockDockerClient) ImagePull(ctx context.Context, ref string, options dockerTypes.ImagePullOptions) (io.ReadCloser, error) {
	return nil, nil
}

func (cli *MockDockerClient) ImagePush(ctx context.Context, ref string, options dockerTypes.ImagePushOptions) (io.ReadCloser, error) {
	return nil, nil
}

func (cli *MockDockerClient) ImageRemove(ctx context.Context, image string, options dockerTypes.ImageRemoveOptions) ([]dockerTypes.ImageDeleteResponseItem, error) {
	return nil, nil
}

func (cli *MockDockerClient) ImageSearch(ctx context.Context, term string, options dockerTypes.ImageSearchOptions) ([]registry.SearchResult, error) {
	return nil, nil
}

func (cli *MockDockerClient) ImageSave(ctx context.Context, images []string) (io.ReadCloser, error) {
	return nil, nil
}

func (cli *MockDockerClient) ImageTag(ctx context.Context, image, ref string) error {
	return nil
}

func (cli *MockDockerClient) ImagesPrune(ctx context.Context, pruneFilter filters.Args) (dockerTypes.ImagesPruneReport, error) {
	return dockerTypes.ImagesPruneReport{}, nil
}
