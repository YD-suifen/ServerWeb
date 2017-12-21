package main

import (
	_ "ServerWeb/routers"
	"github.com/astaxie/beego"
	"ServerWeb/models"
	//"github.com/Unknwon/com"
)

func main() {

    //启动是注册数据库
	models.RegisterDB()

	beego.BConfig.WebConfig.Session.SessionOn = true


	beego.Run()
}

