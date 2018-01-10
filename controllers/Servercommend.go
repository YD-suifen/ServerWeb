package controllers

import (
	"github.com/astaxie/beego"
	"ServerWeb/models"
	"fmt"
	"ServerWeb/sshcommend"

	"ServerWeb/usersessionget"
)

type CommendController struct {
	beego.Controller
}

func (c *CommendController) Commend() {

	a := usersessionget.UserGet(c.Ctx)

	if a == ""{
		c.Redirect("/login", 302)
		return
	}


	c.TplName = "servercommend.html"
}

func (c *CommendController) CommendAction()  {
	a := usersessionget.UserGet(c.Ctx)

	if a == ""{
		c.Redirect("/login", 302)
		return
	}

	ip := c.GetString("ip")
	user := c.GetString("user")
	commend := c.GetString("commend")
    //获取服务器密码
	pwd, err1 := models.SelectServerUserPass(ip)
	fmt.Println(ip,user,commend,pwd)
	if err1 != nil{
		fmt.Println(err1)

		c.Data["commend"] = commend
		c.Data["user"] = user
		c.Data["ip"] = ip
		c.Data["commedinfo"] = "没有此服务器"
		c.TplName = "servercommend.html"
		return
	}
	//在执行远程ssh命令
	content, err  := sshcommend.Action(ip, user,pwd,commend)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(content)



	c.Data["commedinfo"] = content
	c.Data["commend"] = commend
	c.Data["user"] = user
	c.Data["ip"] = ip

	c.TplName = "servercommend.html"


}
