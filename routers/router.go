package routers

import (
	"ServerWeb/controllers"
	"github.com/astaxie/beego"

)





func init() {









    beego.Router("/", &controllers.MainController{}, "*:Index")
	beego.Router("/login", &controllers.LoginController{}, "get,post:Login")
	beego.Router("/registration", &controllers.LoginController{}, "get:Registyemian")
	beego.Router("/regist", &controllers.LoginController{}, "post:Regist")
	beego.Router("/admin/index", &controllers.MainController{}, "*:Home")
	beego.Router("/admin/sercommend", &controllers.ServerController{}, "*:ServerCMD")
	beego.Router("/admin/servercommend", &controllers.CommendController{}, "get:Commend")
	beego.Router("/admin/servercommend", &controllers.CommendController{}, "post:CommendAction")
	beego.Router("/admin/serverlist", &controllers.ServerListController{}, "*:Index")
	beego.Router("/admin/serveradd", &controllers.ServerListController{}, "post:AddServer")
	beego.Router("/admin/serveradd", &controllers.ServerListController{}, "get:AddServer")
}


