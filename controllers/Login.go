package controllers

import (
	"github.com/astaxie/beego"
	"fmt"
	"ServerWeb/models"

)

type LoginController struct {
	beego.Controller

}

func (c *LoginController) Login()  {

	//username := c.Input().Get("username")
	//password := c.Input().Get("password")
	//获取登录输入信息
	username := c.GetString("username")
	password := c.GetString("password")
	//进行数据库查询验证
	yes := models.SelectUser(username,password)

	fmt.Println(username, password)
	if yes {
		c.SetSession("Adminname","hello")

		c.Redirect("/admin/index", 302)
		return
	}

	//if beego.AppConfig.String("username") == username &&
	//	beego.AppConfig.String("password") == password{
	//	c.Ctx.SetCookie("username", username)
	//	c.Ctx.SetCookie("password", password)
	//	c.Redirect("/admin/index", 302)
	//	return
	//}
	c.TplName = "home.html"
}

func (c *LoginController) Registyemian()  {
	c.TplName = "registration.html"
}

func (c *LoginController) Regist() {
	username := c.GetString("username")
	password := c.GetString("password")

	fmt.Println(username, password)
	//注册用户，插入数据
	err, shifou := models.UserRegist(username, password)
	if shifou {
		c.Redirect("/login", 302)
		return

	}
	fmt.Println(err)

	c.TplName = "registration.html"
	return

}

func (c *LoginController) LogOut()  {

	c.DelSession("Adminname")
	c.Redirect("/login", 302)
	return



}
