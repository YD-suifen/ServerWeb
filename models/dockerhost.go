package models

import (
	"time"
	"github.com/astaxie/beego/orm"
	"fmt"
)

type DockerHost struct {
	Id int
	Ip string
	ServerName string
	CreateTime time.Time `orm:"index"`
}


func AddDockerServer(ip, name string) bool {

	orm.Debug = true
	orm.RunSyncdb("default",false,true)
	o := orm.NewOrm()

	fmt.Println(ip,name)
	server := new(DockerHost)
	server.Ip = ip
	server.ServerName = name
	server.CreateTime = time.Now()
	_ , err := o.Insert(server)

	if err != nil {
		fmt.Println("插入有错误：。。。。",err)
		return false
	}
	return true

}

func GetDockerServer() ([]*DockerHost, error)  {
	o := orm.NewOrm()
	cates := make([]*DockerHost,0)
	qs := o.QueryTable("DockerHost")
	_,err := qs.All(&cates)
	fmt.Println("ahixingle ......")
	return cates, err

}

func DelDockerServer(ip string) error {
	o := orm.NewOrm()

	_, err := o.QueryTable(&DockerHost{}).Filter("ip", ip).Delete()
	fmt.Println("caooooo: ",err)

	return err
}