package controllers

import (
	"github.com/astaxie/beego"
	"ServerWeb/models"
	"fmt"
	"ServerWeb/usersessionget"
)

type ServerListController struct {
	beego.Controller
}

func (c *ServerListController) Index()  {
	//a := c.GetSession("yonghu")

	a := usersessionget.UserGet(c.Ctx)

	if a != "" {
		fmt.Println(a)

		var err error
		//渲染页面，查询数据库服务器列表
		c.Data["ServerMach"], err = models.GetServerMach()
		if err != nil {
			fmt.Println(err)
		}

		c.TplName = "serverlist.html"



		return
	}

	c.Redirect("/login", 302)
	return


}

func (c *ServerListController) AddServer()  {


	a := usersessionget.UserGet(c.Ctx)

	if a == ""{
		c.Redirect("/login", 302)
		return
	}
	ip := c.GetString("ip")
	user := c.GetString("user")
	pass := c.GetString("pass")

	if ip != "" && user != "" && pass != "" {
        //添加服务器
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
