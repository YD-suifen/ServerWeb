package models

import (
	"time"
	//"path"
	//"github.com/Unknwon/com"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"

	"fmt"

)


const (
	_SQLITE3_DRIVER = "mysql"
)

func RegisterDB()  {
	//if !com.IsExist(_DB_NAME){
	//	os.MkdirAll(path.Dir(_DB_NAME), os.ModePerm)
	//	os.Create(_DB_NAME)
	//}
	//注册表模型
	orm.RegisterModel(new(User),new(Server))
	//注册数据库驱动
	orm.RegisterDriver(_SQLITE3_DRIVER, orm.DRMySQL)
	//注册数据库连接参数
	orm.RegisterDataBase("default", _SQLITE3_DRIVER,"root:jiange123@/Serverweb?charset=utf8")

}


//func AdminTable()  {
//	orm.Debug = true
//	orm.RunSyncdb("default",false,true)
//	o := orm.NewOrm()
//	o.Using("default")
//	user :=new(User)
//	user.Name = beego.AppConfig.String("username")
//	user.CreateTime = time.Now()
//	user.Pwd = beego.AppConfig.String("password")
//	o.Insert(user)
//}

//创建页面登录用户表
func UserRegist(name, pass string) (error, bool) {

	orm.Debug = true
	//启动自动建表功能
	orm.RunSyncdb("default",false,true)
    //创建一个Ormer对象
	o := orm.NewOrm()
	//初始化表模型，返回一个指针
	user := new(User)
	user.Name = name
	user.Pwd = pass
	user.CreateTime = time.Now()
	//插入数据，
	_ , err := o.Insert(user)
	if err != nil {

		return err,false
	}
	return nil, true

}


//登录页面用户表查询
func SelectUser(name, pass string) bool {
	//声明一个user变量类型为数组，里面的值类型为User类型
	var user []User
	o := orm.NewOrm()
	o.Using("default")
	//Raw查询语句，执行sql，
	_, err := o.Raw("select name,pwd from user where name=?",name).QueryRows(&user)

	fmt.Println(user)

	if err != nil{
		fmt.Println("没有此用户",err,user)
		return false
	}

	for index, _  := range user {
		fmt.Println(user[index].Pwd)

		if user[index].Pwd == pass {
			return true
		}else {
			fmt.Println("dsdsdfsfsd密码错误")
			return false
		}
	}
	return false
}


//服务器密码查询，根据ip字段执行sql查询
func SelectServerUserPass(ip string) (pass string, err error)  {

	var user []Server
	var password string
	o := orm.NewOrm()
	_ , err2 := o.Raw("select pass from server where ip=?", ip).QueryRows(&user)
	if err2 !=nil{
		fmt.Println("cuole................",err2)
		return "", err2
	}

	//user := new(Server)
	//
	//qs := o.QueryTable("server")
	//err2 := qs.Filter("ip", string).One(user)



	hang := user[0]
	password = hang.Pass
	fmt.Println(hang)
	fmt.Println("minahsi:.......",string(password))
	return password, err2





}


