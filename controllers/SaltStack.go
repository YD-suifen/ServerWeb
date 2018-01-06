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
	//"github.com/ghodss/yaml"
	"ServerWeb/usersessionget"
	//"github.com/go-ini/ini"
	"crypto/tls"
)

type Auth_token struct {
	Return []Eauth2 `json:"return"`
}

type Eauth2 struct {
	Eauth string `json:"eauth"`
	Expire float32 `json:"expire"`
	Perms []string `json:"perms"`
	Start float32 `json:"start"`
	Token string `json:"token"`
	User string `json:"user"`
}

type Salt_login struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Eauth  string `json:"eauth"`
}

func Tokend() (string, error)  {

	var requeslogin Salt_login

	saltloginapi := "https://61.147.125.29:8889/login"


	requeslogin.Username = "saltapi"
	requeslogin.Password = "jiange123"
	requeslogin.Eauth = "pam"
	requesjson, err := json.Marshal(requeslogin)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("zheshitokennnnn", string(requesjson))

	requestjsoninfo, err2 := http.NewRequest("POST", saltloginapi, bytes.NewReader(requesjson))
	if err2 != nil {
		fmt.Println(err2)
	}
	requestjsoninfo.Header.Set("Accept", "application/json")
	requestjsoninfo.Header.Set("Content-Type", "application/json")

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: true,
		},
	}


	client := &http.Client{Transport: tr}
	resp, err3 := client.Do(requestjsoninfo)

	//defer resp.Body.Close()
	//fmt.Println(resp.StatusCode)


	if err3 != nil {
		fmt.Println("token huoqu shibai....",err3)
		return "", err3
	}
	body ,err4 := ioutil.ReadAll(resp.Body)
	var jiange Auth_token

	jsonerr := json.Unmarshal([]byte(body), &jiange)
	if jsonerr != nil {
		fmt.Println("jiexi...", jsonerr)
	}
	if err4 != nil {
		fmt.Println("duqu ...",err4)
	}
	fmt.Println(jiange)
	var tocken string
	for _, v := range jiange.Return{
		tocken = v.Token
	}
	fmt.Println("huoqude token is :..", tocken)
	return tocken, nil

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

func Exec_commend(zhuji string, commend string)  CommendRS {

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
	tocken, _ := Tokend()
	requestjsoninfo.Header.Set("X-Auth-Token", tocken)
	requestjsoninfo.Header.Set("Accept", "application/json")

	requestjsoninfo.Header.Set("Content-Type", "application/json")

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: true,
		},
	}
	client := &http.Client{Transport: tr}

	resp, err3 := client.Do(requestjsoninfo)
	if err3 != nil {

		fmt.Println(err3)
	}
	body ,err4 := ioutil.ReadAll(resp.Body)
	var jiange CommendRS

	jsonerr := json.Unmarshal([]byte(body), &jiange)
	if jsonerr != nil {
		fmt.Println(jsonerr)
	}
	if err4 != nil {
		fmt.Println(err4)
	}
	fmt.Println(jiange)

	return jiange


}



type SaltController struct {
	beego.Controller

}

func (c *SaltController) Execution()  {

	a := usersessionget.UserGet(c.Ctx)

	if a == ""{
		c.Redirect("/login", 302)
		return
	}


	c.TplName = "saltremoteexecution.html"
}

func (c *SaltController) ExecutionAction()  {
	a := usersessionget.UserGet(c.Ctx)

	if a == ""{
		c.Redirect("/login", 302)
		return
	}


	zhujixinxi := c.GetString("zhujixinxi")
	commend := c.GetString("commend")
	fmt.Println(zhujixinxi,commend)

	jieguo:= Exec_commend(zhujixinxi, commend)


	var cache bytes.Buffer
	for _, i := range jieguo.Return{
		for key, val := range i {
			cache.WriteString("\n-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-\n")

			cache.WriteString(fmt.Sprint("主机：\n",key))

			cache.WriteString(val)
			cache.WriteString("\n-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-\n")

		}
	}
	date := cache.String()

	fmt.Println("sdfasdfasdfsadfsdf", date)

	c.Data["zhujixinxi"] = zhujixinxi
	c.Data["commend"] = commend
	c.Data["commedinfo"] = date

	c.TplName = "saltremoteexecution.html"
}






