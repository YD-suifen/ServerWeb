package controllers

import (
	"github.com/astaxie/beego"
	"ServerWeb/models"
	"fmt"
	"ServerWeb/sshcommend"

)

type CommendController struct {
	beego.Controller
}

func (c *CommendController) Commend() {


	c.TplName = "servercommend.html"
}

func (c *CommendController) CommendAction()  {


	ip := c.GetString("ip")
	user := c.GetString("user")
	commend := c.GetString("commend")

	pwd, err1 := models.SelectServerUserPass(ip)
	fmt.Println(ip,user,commend,pwd)
	if err1 !=nil{
		fmt.Println(err1)
		return
	}
	content, err  := sshcommend.Action(ip, user,pwd,commend)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(content)



	c.Data["commedinfo"] = content
	c.TplName = "servercommend.html"


}
