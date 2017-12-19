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
	orm.RegisterModel(new(User),new(Server))
	orm.RegisterDriver(_SQLITE3_DRIVER, orm.DRMySQL)
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

func UserRegist(name, pass string) (error, bool) {
	orm.Debug = true
	orm.RunSyncdb("default",false,true)
	o := orm.NewOrm()
	user := new(User)
	user.Name = name
	user.Pwd = pass
	user.CreateTime = time.Now()
	_ , err := o.Insert(user)
	if err != nil {

		return err,false
	}
	return nil, true

}

func SelectUser(name, pass string) bool {
	var user []User
	o := orm.NewOrm()
	o.Using("default")
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


