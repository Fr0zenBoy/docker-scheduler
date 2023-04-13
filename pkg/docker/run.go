package docker

import (
	"context"
	"io"
	"os"

	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
)

func imageExist(ctx context.Context, cli *client.Client, config container.Config) (bool, string){
	images, err := cli.ImageList(ctx, types.ImageListOptions{})
	if err != nil {
		panic(err)
	}

	//TODO pass the container name without tag
	var exist bool
	var imageId string
	for _, image := range images {
		for _, tags := range image.RepoTags {
			if config.Image == tags {
				exist = true
				imageId = image.ID
				break
			}
		}
	}
	return exist, imageId
}

func pullContainer(ctx context.Context, cli *client.Client, config container.Config) io.ReadCloser {

	//TODO add the option of other repositories
	out, err := cli.ImagePull(ctx, config.Image, types.ImagePullOptions{})
	if err != nil {
		panic(err)
	}
	defer out.Close()

	return out
}

func build(ctx context.Context, cli *client.Client, config container.Config, containerName string) container.CreateResponse {

	resp, err := cli.ContainerCreate(ctx, &config, nil, nil, nil, containerName)
	if err != nil {
		panic(err)
	}

	return resp
}

func RunContainer(ctx context.Context, cli *client.Client, config container.Config, containerName string) error {

	if exist, _ := imageExist(ctx, cli, config); !exist {
		out := pullContainer(ctx, cli, config)
		io.Copy(os.Stdout, out)
	}

	resp := build(ctx, cli, config, containerName) 

	err := cli.ContainerStart(ctx, resp.ID, types.ContainerStartOptions{})
	if err != nil {
		panic(err)
	}

	return err

}
