package controllers

import (
	"github.com/astaxie/beego"

	"fmt"
	//
	//"github.com/docker/docker/api/types"
	//"github.com/docker/docker/client"
	//"golang.org/x/net/context"
	//"net/http"
	//"image/draw"
	"ServerWeb/usersessionget"
	"ServerWeb/models"
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
