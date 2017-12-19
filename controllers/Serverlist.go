package controllers

import (
	"github.com/astaxie/beego"
	"ServerWeb/models"
	"fmt"
)

type ServerListController struct {
	beego.Controller
}

func (c *ServerListController) Index()  {

	var err error

	c.Data["ServerMach"], err = models.GetServerMach()
	if err != nil {
		fmt.Println(err)
	}

	c.TplName = "serverlist.html"
}

func (c *ServerListController) AddServer()  {
	ip := c.GetString("ip")
	user := c.GetString("user")
	pass := c.GetString("pass")

	if ip != "" && user != "" && pass != "" {

		shi := models.AddServerMach(ip,user,pass)
		fmt.Println(shi)
		if shi {

			c.Redirect("/admin/serverlist",302)
			return

		}

	}

	c.TplName = "serveradd.html"
	return
}
