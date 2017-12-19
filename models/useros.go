package models

import (
	"time"
	//"path"
	//"github.com/Unknwon/com"
	//"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	//"github.com/astaxie/beego"



)

type User struct {
	Id int
	Name string
	Pwd string
	CreateTime time.Time `orm:"index"`
}





