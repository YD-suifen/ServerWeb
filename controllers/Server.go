package controllers


import (
	"github.com/astaxie/beego"
)

type ServerController struct {
	beego.Controller
}

func (c *ServerController) ServerCMD() {



	c.TplName= "sercommend.html"
}
