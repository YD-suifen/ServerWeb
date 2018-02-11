package docker


import (
	"fmt"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
	"golang.org/x/net/context"
	"net/http"
)

func Containers(host string) []string {

	var cli *http.Client
	host := "http://133.130.122.48:2378"
	version := "v1.35"
	ctx, err := client.NewClient(host, version, cli, nil)
	if err != nil {
		panic(err)
	}

	containers, err := ctx.ContainerList(context.Background(), types.ContainerListOptions{})

	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	for _, container := range containers {
		fmt.Println("containers list :",container.ID)
	}


}