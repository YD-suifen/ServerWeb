package controllers

import (
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Index() {



	//c.TplName = "index.tpl"
	c.TplName = "home.html"
}

func (c *MainController) Home()  {


	c.TplName = "zhuye.html"
}
