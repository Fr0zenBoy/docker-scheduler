package container

import (
	"context"
	"io"
	"os"

	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
)

type DockerContainer struct {

	containerName string
	ctx context.Context
	client client.Client
	config container.Config
}

func (d DockerContainer) imageExist() (bool, string){
	images, err := d.client.ImageList(d.ctx, types.ImageListOptions{})
	if err != nil {
		panic(err)
	}

	//TODO pass the container name without tag
	var exist bool
	var imageId string
	for _, image := range images {
		for _, tags := range image.RepoTags {
			if d.config.Image == tags {
				exist = true
				imageId = image.ID
				break
			}
		}
	}
	return exist, imageId
}

func (d DockerContainer) pullContainer() io.ReadCloser {

	//TODO add the option of other repositories
	out, err := d.client.ImagePull(d.ctx, d.config.Image, types.ImagePullOptions{})
	if err != nil {
		panic(err)
	}
	defer out.Close()

	return out
}

func (d DockerContainer) build() container.CreateResponse {

	resp, err := d.client.ContainerCreate(d.ctx, &d.config, nil, nil, nil, d.containerName)
	if err != nil {
		panic(err)
	}

	return resp
}

func (d DockerContainer) RunContainer() error {

	if exist, _ := d.imageExist(); !exist {
		out := d.pullContainer()
		io.Copy(os.Stdout, out)
	}

	resp := d.build() 

	err := d.client.ContainerStart(d.ctx, resp.ID, types.ContainerStartOptions{})
	if err != nil {
		panic(err)
	}

	return err

}

func NewContainer(imageName, containerName string, containerConfig container.Config) DockerContainer {

	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		panic(err)
	}
	defer cli.Close()

	dc := DockerContainer {
		containerName: containerName,
		ctx: context.Background(),
		client: *cli,
		}

	dc.config.Image = imageName

	return dc
}
