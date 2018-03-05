package routers

import (
	"ServerWeb/controllers"
	"github.com/astaxie/beego"

)





func init() {



    beego.Router("/", &controllers.MainController{}, "*:Index")
	beego.Router("/login", &controllers.LoginController{}, "get,post:Login")
	beego.Router("/logout", &controllers.LoginController{}, "get,post:LogOut")
	beego.Router("/registration", &controllers.LoginController{}, "get:Registyemian")
	beego.Router("/regist", &controllers.LoginController{}, "post:Regist")
	beego.Router("/admin/index", &controllers.MainController{}, "*:Home")
	beego.Router("/admin/sercommend", &controllers.ServerController{}, "*:ServerCMD")
	beego.Router("/admin/servercommend", &controllers.CommendController{}, "get:Commend")
	beego.Router("/admin/servercommend", &controllers.CommendController{}, "post:CommendAction")
	beego.Router("/admin/serverlist", &controllers.ServerListController{}, "*:Index")
	beego.Router("/admin/serveradd", &controllers.ServerListController{}, "post,get:AddServer")
	//beego.Router("/admin/serveradd", &controllers.ServerListController{}, "get:AddServer")
	beego.Router("/admin/saltremoteexecution", &controllers.SaltController{}, "get:Execution")
	beego.Router("/admin/saltremoteexecution", &controllers.SaltController{}, "post:ExecutionAction")
	beego.Router("/admin/saltkeylist", &controllers.SaltController{}, "post,get:KeyListAllAction")
	beego.Router("/admin/saltfileCp", &controllers.CpfileController{}, "post:CpfileAction")
	beego.Router("/admin/saltfileCp", &controllers.CpfileController{}, "get:CpfileGet")
	beego.Router("/admin/updatafile", &controllers.CpfileController{}, "post:UpDataFile")
	beego.Router("/admin/docker", &controllers.DockerController{}, "post,get:Action")
	beego.Router("/admin/dockerhostinfo", &controllers.DockerController{}, "post,get:Containers")

}


