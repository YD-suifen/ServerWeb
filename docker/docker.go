package docker


import (


	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
	"golang.org/x/net/context"
	"net/http"
)

func AllContainers(host string) ([]types.Container, error){

	var cli *http.Client
	hosturl := "http://" + host
	version := "v1.35"
	ctx, err := client.NewClient(hosturl, version, cli, nil)
	if err != nil {
		return nil, err
	}

	containers, err := ctx.ContainerList(context.Background(), types.ContainerListOptions{})

	if err != nil {
		return nil, err
	}


	return containers, nil

}

func AllImages(host string) ([]types.ImageSummary, error){
	var cli *http.Client
	hosturl := "http://" + host
	version := "v1.35"
	ctx, err := client.NewClient(hosturl, version, cli, nil)
	if err != nil {
		return nil, err
	}
	images, err := ctx.ImageList(context.Background(), types.ImageListOptions{})
	if err != nil {
		return nil, err
	}


	return images, nil
}

func AllNetworkMode(host string) ([]types.NetworkResource, error) {
	var cli *http.Client
	hosturl := "http://" + host
	version := "v1.35"
	ctx, err := client.NewClient(hosturl, version, cli, nil)
	if err != nil {
		return nil, err
	}
	networkmode, err := ctx.NetworkList(context.Background(), types.NetworkListOptions{})
	if err != nil {
		return nil, err
	}


	return networkmode, nil
}