package controllers

import (
	"github.com/astaxie/beego"

	"fmt"
	//"github.com/docker/docker/api/types"

	"ServerWeb/docker"
	"ServerWeb/usersessionget"
	"ServerWeb/models"
	"strings"
)

type DockerController struct {
	beego.Controller

}



func (c *DockerController) Action()  {

	a := usersessionget.UserGet(c.Ctx)

	if a == ""{
		c.Redirect("/login", 302)
		return
	}

	var nodes []string
	ipport := c.GetString("ip")
	dockername := c.GetString("dockername")
	fmt.Println("this is 22222:", ipport, dockername)
	op := c.GetString("op")

	switch op {
	case "add":
		nodes = append(nodes,ipport)
		what := models.AddDockerServer(ipport, dockername)
		if what {
			c.Redirect("/admin/docker",301)
			return

		}
	case "del":
		ip := c.Input().Get("ip")

		fmt.Println("delete",ip)


		err := models.DelDockerServer(ip)
		fmt.Println("this delete is ",err)
		if err != nil {
			fmt.Println("error delete data")
			return
		}
		c.Redirect("/admin/docker",302)
		return

	}

	dockerhost, err := models.GetDockerServer()
	if err != nil {
		fmt.Println(err)
	}

	c.Data["DockerHost"] = dockerhost

	c.TplName = "docker.html"

}

func (c *DockerController) Containers() {

	a := usersessionget.UserGet(c.Ctx)
	if a == ""{
		c.Redirect("/login", 302)
		return
	}


	containerdel := c.Input().Get("op")

	if containerdel != ""  {
		host := c.Input().Get("ip")

		switch containerdel {
		case "containerdel":
			containerid := c.Input().Get("id")
			what, _ := docker.Containerfalesdelete(host,containerid)

			if what {

				c.Redirect("/admin/dockerhostinfo.html?ip="+host, 301)
				return
			}
		case "imagedel":
			imagesid := c.Input().Get("id")
			what := docker.Imagesdelete(host,imagesid)
			if what {
				c.Redirect("/admin/dockerhostinfo.html?ip="+host, 302)
				return
			}
		case "networkdel":
			networkid := c.Input().Get("id")
			what := docker.NetworkModedelete(host, networkid)
			if what {
				c.Redirect("/admin/dockerhostinfo.html?ip="+host, 303)
			}

		}
	}


	ip := c.Input().Get("ip")


	containers, err := docker.AllContainers(ip)

	images, err := docker.AllImages(ip)

	networkmode, err := docker.AllNetworkMode(ip)

	if err != nil {
		return
	}
	for _, k := range containers{
		fmt.Println("zhe shi kong zhiqi hanshu :", k.Container.ID)
	}


	strings.Contains("widuu", "wi")

	c.Data["Containers"] = containers
	c.Data["Images"] = images
	c.Data["NetworkModes"] = networkmode
	c.TplName = "dockerhostinfo.html"

}
