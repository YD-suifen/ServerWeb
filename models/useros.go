package models

import (
	"time"
	//"path"
	//"github.com/Unknwon/com"
	//"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	//"github.com/astaxie/beego"



)


//登录页面用户表结构
type User struct {
	Id int
	Name string
	Pwd string
	CreateTime time.Time `orm:"index"`
}





