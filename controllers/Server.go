package controllers


import (
	"github.com/astaxie/beego"
	"ServerWeb/usersessionget"
)

type ServerController struct {
	beego.Controller
}

func (c *ServerController) ServerCMD() {
	a := usersessionget.UserGet(c.Ctx)

	if a == ""{
		c.Redirect("/login", 302)
		return
	}



	c.TplName= "sercommend.html"
}
