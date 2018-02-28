package models

import (
	"time"

	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"github.com/astaxie/beego"

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
	MysqlUser := beego.AppConfig.String("mysqluser")
	MysqlPass := beego.AppConfig.String("mysqlpass")
	MysqlHost := beego.AppConfig.String("mysqlhost")
	orm.RegisterModel(new(User),new(Server),new(DockerHost))
	//注册数据库驱动
	orm.RegisterDriver(_SQLITE3_DRIVER, orm.DRMySQL)
	//注册数据库连接参数
	//orm.RegisterDataBase("default", _SQLITE3_DRIVER,"root:jiange123@/Serverweb?charset=utf8")
	orm.RegisterDataBase("default", _SQLITE3_DRIVER,MysqlUser+":"+MysqlPass+"@tcp(" + MysqlHost + ")/Serverweb?charset=utf8")

}


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

	var user Server

	o := orm.NewOrm()
	err1 := o.Raw("SELECT * FROM server WHERE ip = ?",ip).QueryRow(&user)
	if err1 == orm.ErrNoRows {

		fmt.Println("meiyou 服务器", err1)
		return "", err1
	}


	return user.Pass, nil


}


