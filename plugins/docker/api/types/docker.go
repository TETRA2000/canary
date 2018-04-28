package types

import (
	"github.com/docker/docker/client"
	"github.com/docker/docker/api/types"
	"golang.org/x/net/context"
	"io"
	"github.com/docker/docker/api/types/image"
	"github.com/docker/docker/api/types/registry"
	"github.com/docker/docker/api/types/filters"
)

type IDockerClient interface {
	client.ImageAPIClient
}

type DockerClient struct {
	RealClient *client.Client
	IDockerClient
}

// TODO find efficient way to delegate method call.

func (cli *DockerClient) ImageBuild(ctx context.Context, context io.Reader, options types.ImageBuildOptions) (types.ImageBuildResponse, error) {
	return cli.RealClient.ImageBuild(ctx, context, options)
}

func (cli *DockerClient) BuildCachePrune(ctx context.Context) (*types.BuildCachePruneReport, error) {
	return cli.RealClient.BuildCachePrune(ctx)
}

func (cli *DockerClient) ImageCreate(ctx context.Context, parentReference string, options types.ImageCreateOptions) (io.ReadCloser, error){
	return cli.RealClient.ImageCreate(ctx, parentReference, options)
}

func (cli *DockerClient) ImageHistory(ctx context.Context, image string) ([]image.HistoryResponseItem, error) {
	return cli.RealClient.ImageHistory(ctx, image)
}

func (cli *DockerClient) ImageImport(ctx context.Context, source types.ImageImportSource, ref string, options types.ImageImportOptions) (io.ReadCloser, error) {
	return cli.RealClient.ImageImport(ctx, source, ref, options)
}

func (cli *DockerClient) ImageInspectWithRaw(ctx context.Context, image string) (types.ImageInspect, []byte, error) {
	return cli.RealClient.ImageInspectWithRaw(ctx, image)
}

func (cli *DockerClient) ImageList(ctx context.Context, options types.ImageListOptions) ([]types.ImageSummary, error) {
	return cli.RealClient.ImageList(ctx, options)
}

func (cli *DockerClient) ImageLoad(ctx context.Context, input io.Reader, quiet bool) (types.ImageLoadResponse, error) {
	return cli.RealClient.ImageLoad(ctx, input, quiet)
}

func (cli *DockerClient) ImagePull(ctx context.Context, ref string, options types.ImagePullOptions) (io.ReadCloser, error) {
	return cli.RealClient.ImagePull(ctx, ref, options)
}

func (cli *DockerClient) ImagePush(ctx context.Context, ref string, options types.ImagePushOptions) (io.ReadCloser, error) {
	return cli.RealClient.ImagePush(ctx, ref, options)
}

func (cli *DockerClient) ImageRemove(ctx context.Context, image string, options types.ImageRemoveOptions) ([]types.ImageDeleteResponseItem, error) {
	return cli.RealClient.ImageRemove(ctx, image, options)
}

func (cli *DockerClient) ImageSearch(ctx context.Context, term string, options types.ImageSearchOptions) ([]registry.SearchResult, error) {
	return cli.RealClient.ImageSearch(ctx, term, options)
}

func (cli *DockerClient) ImageSave(ctx context.Context, images []string) (io.ReadCloser, error) {
	return cli.RealClient.ImageSave(ctx, images)
}

func (cli *DockerClient) ImageTag(ctx context.Context, image, ref string) error {
	return cli.RealClient.ImageTag(ctx, image, ref)
}

func (cli *DockerClient) ImagesPrune(ctx context.Context, pruneFilter filters.Args) (types.ImagesPruneReport, error) {
	return cli.RealClient.ImagesPrune(ctx, pruneFilter)
}

