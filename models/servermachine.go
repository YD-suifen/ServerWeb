package models

import (
	"time"
	"github.com/astaxie/beego/orm"
	"fmt"
)

//服务器表结构
type Server struct {
	Id int
	Ip string
	Name string
	Pass string
	CreateTime time.Time `orm:"index"`
}

//添加服务器列表
func AddServerMach(ip, user, pass string) bool {

	orm.Debug = true
	orm.RunSyncdb("default",false,true)
	o := orm.NewOrm()

	fmt.Println(ip,user,pass)
	server := new(Server)
	server.Ip = ip
	server.Name = user
	server.Pass= pass
	server.CreateTime = time.Now()
	_ , err := o.Insert(server)

	if err != nil {
		fmt.Println("插入有错误：。。。。",err)
		return false
	}
	return true

}

//func GetServerMach()  {
//
//	orm.Debug = true
//	o := orm.NewOrm()
//
//	server := make([]*ServerMach,0)
//	qs := o.QueryTable("ServerMach")
//	_, err2 = qs.OrderBy("-created").All(&topics)
//	}else {
//		_, err2 = qs.All(&topics)
//	}
//
//	return topics, err2
//
//
//}


//获取数据库内，服务器表的所有数据，服务器列表
func GetServerMach() ([]*Server, error)  {
	o := orm.NewOrm()
	cates := make([]*Server,0)
	qs := o.QueryTable("Server")
	_,err := qs.All(&cates)
	fmt.Println("ahixingle ......")
	return cates, err

}

