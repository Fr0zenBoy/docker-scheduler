package container

import (
	"context"
	"fmt"
	"io"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
)

func ImageExist(ctx context.Context, cli client.Client, imageName string) (bool, string){
	images, err := cli.ImageList(ctx, types.ImageListOptions{})
	if err != nil {
		panic(err)
	}

	//TODO pass the container name without tag
	var exist bool
	var imageID string
	for _, image := range images {
		for _, tags := range image.RepoTags {
			if imageName == tags {
				exist = true
				imageID = image.ID
				break
			}
		}
	}
	return exist, imageID
}

func PullContainer(ctx context.Context, cli client.Client, imageName string) io.ReadCloser {

	//TODO add the option of other repositories
	out, err := cli.ImagePull(ctx, imageName, types.ImagePullOptions{})
	if err != nil {
		panic(err)
	}

	return out
}
