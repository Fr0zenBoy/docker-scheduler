package container

import (
	"context"
	"io"

	"github.com/docker/docker/client"
	"github.com/docker/docker/api/types"
)

func PullContainer(ctx context.Context, cli client.Client, imageName string) io.ReadCloser {

	//TODO add the option of other repositories
	out, err := cli.ImagePull(ctx, imageName, types.ImagePullOptions{})
	if err != nil {
		panic(err)
	}

	return out
}
