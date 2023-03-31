package container

import (
	"context"
	"fmt"
	"io"
	"os"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
)

func RunContainer(imageName string) error {
	ctx := context.Background()
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		panic(err)
	}
	defer cli.Close()

	if exist, _ := ImageExist(ctx, *cli, imageName); !exist {
		out := PullContainer(ctx, *cli, imageName)
		defer out.Close()
		io.Copy(os.Stdout, out)
	}

	resp := Build(ctx, *cli, imageName) 

	if err := cli.ContainerStart(ctx, resp.ID, types.ContainerStartOptions{}); err != nil {
		panic(err)
	}

	fmt.Println(resp.ID)

	return err

}
