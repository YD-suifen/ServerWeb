package main

import (
	_ "ServerWeb/routers"
	"github.com/astaxie/beego"
	"ServerWeb/models"
	//"github.com/Unknwon/com"
)

func main() {

	models.RegisterDB()

	beego.BConfig.WebConfig.Session.SessionOn = true


	beego.Run()
}

