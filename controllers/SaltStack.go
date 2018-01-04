package controllers

import (
	"github.com/astaxie/beego"
	//"fmt"
	//"ServerWeb/models"
	"encoding/json"
	"net/http"
	//
	//"bytes"
	//"io/ioutil"
	"fmt"
	"bytes"
	"io/ioutil"
	"github.com/ghodss/yaml"
)

type Auth_token struct {
	Return []Eauth2 `json:"return"`
}

type Eauth2 struct {
	Eauth string `json:"eauth"`
	Expire string `json:"expire"`
	Perms []string `json:"perms"`
	Start string `json:"start"`
	Tocken string `json:"tocken"`
	User string `json:"user"`
}

type Salt_login struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Eauth  string `json:"eauth"`
}

func Tokend() string  {

	var requeslogin Salt_login

	saltapi := beego.AppConfig.String("saltloginapi")
	saltapiname := beego.AppConfig.String("saltapi")
	saltapipass := beego.AppConfig.String("jiange123")
	requeslogin.Username = saltapiname
	requeslogin.Password = saltapipass
	requeslogin.Eauth = "pam"
	requesjson, err := json.Marshal(requeslogin)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(requesjson)

	requestjsoninfo, err2 := http.NewRequest("POST", saltapi, bytes.NewReader(requesjson))
	if err2 != nil {
		fmt.Println(err2)
	}
	requestjsoninfo.Header.Set("Accept", "application/json")

	requestjsoninfo.Header.Set("Content-Type", "application/json")
	client := http.Client{}
	resp, err3 := client.Do(requestjsoninfo)
	if err3 != nil {
		fmt.Println(err3)
	}
	body ,err4 := ioutil.ReadAll(resp.Body)
	var jiange Auth_token

	jsonerr := json.Unmarshal([]byte(body), &jiange)
	if jsonerr != nil {
		fmt.Println(jsonerr)
	}
	if err4 != nil {
		fmt.Println(err4)
	}
	fmt.Println(jiange)
	var tocken string
	for _, v := range jiange.Return{
		tocken = v.Tocken
	}
	return tocken

}

type ActionCommend struct {
	Client string `json:"client"`
	Tgt string `json:"tgt"`
	Fun string `json:"fun"`
	Arg string `json:"arg"`
}

type CommendRS struct {
	Return []map[string]string `json:"return"`
} 

func exec_commend(zhuji string, commend string)  string {

	var action ActionCommend
	action.Client = "local"
	action.Tgt = zhuji
	action.Fun = "cmd.run"
	action.Arg = commend

	actionjson, err := json.Marshal(action)
	if err != nil {
		fmt.Println(err)
	}
	saltapi := beego.AppConfig.String("saltapi")

	requestjsoninfo, err2 := http.NewRequest("POST", saltapi, bytes.NewReader(actionjson))
	if err2 != nil {
		fmt.Println(err2)
	}
	tocken := Tokend()
	requestjsoninfo.Header.Set("X-Auth-Token", tocken)
	requestjsoninfo.Header.Set("Accept", "application/json")

	requestjsoninfo.Header.Set("Content-Type", "application/json")
	client := http.Client{}
	resp, err3 := client.Do(requestjsoninfo)
	if err3 != nil {
		fmt.Println(err3)
	}
	body ,err4 := ioutil.ReadAll(resp.Body)
	var jiange CommendRS

	jsonerr := json.Unmarshal([]byte(body), &CommendRS)
	if jsonerr != nil {
		fmt.Println(jsonerr)
	}
	if err4 != nil {
		fmt.Println(err4)
	}
	fmt.Println(jiange)

	xml , err8 := yaml.JSONToYAML([]byte(body))
	if err8 !=nil {
		fmt.Println(err8)
	}
	return string(xml)

	//for _, v := range jiange.Return {
	//	if v.(string){
	//		for
	//	}
	//}

}



type SaltController struct {
	beego.Controller

}

func (c *SaltController) Execution()  {

	zhujixinxi := c.Input().Get("zhujixinxi")
	commend := c.Input().Get("commend")

	jieguo := exec_commend(zhujixinxi, commend)

	c.Data["jieguo"] = jieguo

	c.TplName = "saltremoteexecution.html"
}






