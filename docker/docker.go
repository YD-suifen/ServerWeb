package docker


import (


	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
	"golang.org/x/net/context"
	"net/http"
	"fmt"
	"strings"
)

type DockerHostAPI struct {
	types.Container
	types.ImageSummary
	types.NetworkResource
	StatusN bool
	Ip string
}

func AllContainers(host string) ([]DockerHostAPI,error){

	var cli *http.Client
	hosturl := "http://" + host
	version := "v1.35"
	ctx, err := client.NewClient(hosturl, version, cli, nil)
	if err != nil {
		return nil, err
	}

	containers, err := ctx.ContainerList(context.Background(), types.ContainerListOptions{All:true})

	if err != nil {
		return nil, err
	}

	var jieguo2 []DockerHostAPI

	for _, v := range containers{
		fmt.Println("the container status is : ",v.Status)
		if strings.Contains(v.Status, "Up") {
			jieguo := DockerHostAPI{Container:v,StatusN:true, Ip:host}
			jieguo2 = append(jieguo2,jieguo)
		}else {
			jieguo := DockerHostAPI{Container:v,StatusN:false, Ip:host}
			jieguo2 = append(jieguo2,jieguo)
		}

	}

	for _, k := range jieguo2{
		fmt.Println("zhixing kanshu : ", k.Container.ID)
	}

	return jieguo2, nil

}

func AllImages(host string) ([]DockerHostAPI, error){
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
	var jiange []DockerHostAPI
	for _, v := range images{
		jianjie := DockerHostAPI{ImageSummary:v, Ip:host}
		jiange = append(jiange,jianjie)
	}

	fmt.Println("images host2 ip is : ", host)

	for _, k := range jiange{
		fmt.Println("zhixing kanshu : ", k.ImageSummary.ID)
	}


	return jiange, nil
}

func AllNetworkMode(host string) ([]DockerHostAPI, error) {
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

	var jiange []DockerHostAPI
	for _, v := range networkmode{
		jianjie := DockerHostAPI{NetworkResource:v, Ip:host}
		jiange = append(jiange,jianjie)
	}


	return jiange, nil
}


func Containerfalesdelete(host, containerID string) (bool, error) {

	var cli *http.Client
	hosturl := "http://" + host
	version := "v1.35"

	ctx, err := client.NewClient(hosturl, version, cli, nil)
	if err != nil {
		return false, err
	}

	err1 := ctx.ContainerRemove(context.Background(), containerID,types.ContainerRemoveOptions{})

	if err1 != nil {
		return false, err1
	}

	return true, nil

}

func Imagesdelete(host, imageid string) (bool) {
	var cli *http.Client
	hosturl := "http://" + host
	version := "v1.35"

	ctx, err := client.NewClient(hosturl, version, cli,nil)
	if err != nil {
		return false
	}

	_, err1 := ctx.ImageRemove(context.Background(), imageid, types.ImageRemoveOptions{})
	if err1 != nil{
		return false
	}

	return true

}

func NetworkModedelete(host, networkid string) bool {
	var cli *http.Client
	hosturl := "http://" + host
	version := "v1.35"
	ctx, err := client.NewClient(hosturl,version,cli,nil)
	if err != nil {
		return false
	}
	err1 := ctx.NetworkRemove(context.Background(), networkid)
	if err1 != nil {
		return false
	}
	return true

}