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
	"github.com/pkg/errors"
	"time"

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

	//saltloginapi := "https://61.147.125.29:8889/login"
	saltloginapi := beego.AppConfig.String("saltloginapi")
	saltapiname := beego.AppConfig.String("saltapiname")
	saltapipass := beego.AppConfig.String("saltapipass")

	requeslogin.Username = saltapiname
	requeslogin.Password = saltapipass
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


	client := &http.Client{
		Timeout: 5 * time.Second,
		Transport: tr,
		}
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
	Match string `json:"match"`
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
	client := &http.Client{
		Timeout: 5 * time.Second,
		Transport: tr,
		}

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

type ListKey struct {
	Return []struct{
		Tag string `json:"tag"`
		Data struct{
			Jid string `json:"jid"`
			Return struct{
				Minions_pre []interface{} `json:"minions_pre"`
				Minions_rejected []interface{} `json:"minions_rejected"`
				Minions_denied []interface{} `json:"minions_denied"`
				Local []string `json:"local"`
				Minions []string `json:"minions"`
			} `json:"return"`
			Success bool
			Stamp string `json:"_stamp"`
			Tag string `json:"tag"`
			User string `json:"user"`
			Fun string `json:"fun"`

		} `json:"data"`
	} `json:"return"`
}



func KeyListAll() ([]string, error) {
	var action ActionCommend
	action.Client = "wheel"
	action.Tgt = "*"
	action.Fun = "key.list_all"
	actionjson, err := json.Marshal(action)
	if err != nil {
		fmt.Println(err)
	}
	saltapi := beego.AppConfig.String("saltapi")
	requestjsoninfo, err2 := http.NewRequest("POST", saltapi, bytes.NewReader(actionjson))
	if err2 != nil {
		fmt.Println(err2)
	}
	tocken, xinxi := Tokend()
	if xinxi != nil {
		return nil, xinxi
	}
	requestjsoninfo.Header.Set("X-Auth-Token", tocken)
	requestjsoninfo.Header.Set("Accept", "application/json")

	requestjsoninfo.Header.Set("Content-Type", "application/json")

	tr := &http.Transport{

		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: true,
		},
	}
	client := &http.Client{
		Timeout: 5 * time.Second,
		Transport: tr,
		}

	resp, err3 := client.Do(requestjsoninfo)
	if err3 != nil {

		fmt.Println("http baocuo ...", err3)
	}
	body ,err4 := ioutil.ReadAll(resp.Body)
	if err4 != nil {
		fmt.Println(err4)
	}

	var jiange ListKey

	jsonerr := json.Unmarshal([]byte(body), &jiange)
	if jsonerr != nil {
		fmt.Println(jsonerr)
	}

	for _, v := range jiange.Return{

		aa := v.Data.Return.Minions
		return aa, nil


	}
	return nil, nil

}

//func (c *SaltController) KeyList() {
//	a := usersessionget.UserGet(c.Ctx)
//	if a == ""{
//		c.Redirect("/login", 302)
//		return
//	}
//
//
//	list := KeyListAll()
//
//	c.Data["keylist"] = list
//	c.TplName = "saltkeylist.html"
//	return
//
//}


func (c *SaltController) KeyListAllAction()  {

	a := usersessionget.UserGet(c.Ctx)
	if a == ""{
		c.Redirect("/login", 302)
		return
	}



	op := c.Input().Get("op")
	fmt.Println("zhua ququuququu 000", op)


	switch op{
	case "add":
		minionkey := c.GetString("minionkey")
		data, err := KeyAccepted(minionkey)

		fmt.Println("shujuwei 000",data)

		fmt.Println("shujubaocuowei 000", err)

		if err == nil {

			//list := KeyListAll()
			//
			//c.Data["keylist"] = list
			//c.TplName = "saltkeylist.html"
			//return
			c.Redirect("/admin/saltkeylist.html",301)
			return

		}
	case "del":
		ip := c.Input().Get("ip")
		what := KeyDeletAction(ip)

		fmt.Println(what)

		c.Redirect("/admin/saltkeylist.html",302)
		return


	}

	//add := c.GetString("op")
	//if add != "" { //添加逻辑
	//	minionkey := c.GetString("minionkey")
	//	data, err := KeyAccepted(minionkey)
	//
	//	fmt.Println("shujuwei 000",data)
	//
	//	fmt.Println("shujubaocuowei 000", err)
	//
	//	if err == nil {
	//
	//		list := KeyListAll()
	//
	//		c.Data["keylist"] = list
	//		c.TplName = "saltkeylist.html"
	//		return
	//	}
	//	fmt.Println("piao;aing------")
	//
	//	//c.Data["keylist"] = data
	//
	//	c.TplName = "saltkeylist.html"
	//
	//	return
	//}

	//list := KeyListAll()
	//var cache bytes.Buffer
	//for _, i := range list{
	//	cache.WriteString("Minion-\n")
	//	cache.WriteString(i)
	//}
	//date := cache.String()
	//fmt.Println(date)
	//var aa []string
	//aa = append(aa, "1.1.1.1")
	//aa = append(aa, "2.2.2.2")
	//aa = append(aa, "3.3.3.3")
	nihao, err := KeyListAll()
	if err != nil {
		var list []string
		list = append(list,"")

		c.Data["keylist"] = list
		c.TplName = "saltkeylist.html"
		return
	}


	c.Data["keylist"] = nihao
	c.TplName = "saltkeylist.html"
	return


}

type KeyAcceptedST struct {
	Return []struct{
		Tag string `json:"tag"`
		Data struct{
			Jid string `json:"jid"`
			Return struct{
				Minions []string `json:"minions"`
			} `json:"return"`
			Success bool `json:"success"`
			Stamp string `json:"_stamp"`
			Tag string `json:"tag"`
			User string `json:"user"`
			Fun string `json:"fun"`
		} `json:"data"`

	} `json:"return"`
}


type KeyDeletST struct {
	Return []struct{
		Tag string `json:"tag"`
		Data struct{
			Jid string `json:"jid"`
			Return struct{} `json:"return"`
			Success bool `json:"success"`
			Stamp string `json:"_stamp"`
			Tag string `json:"tag"`
			User string `json:"user"`
			Fun string `json:"fun"`
		} `json:"data"`

	} `json:"return"`
}

//func (c *SaltController) Accepted()  {
//
//	a := usersessionget.UserGet(c.Ctx)
//
//	if a == ""{
//		c.Redirect("/login", 302)
//		return
//	}
//	add := c.GetString("add")
//	ip := c.GetString("ip")
//	fmt.Println(add,ip)
//
//	//data := KeyAccepted(minionkey)
//
//
//	//c.Data[""]
//
//	c.TplName = "saltkeylist.html"
//
//
//
//
//}

func KeyAccepted(minionkey string) ([]string, error) {

	var keyjson  ActionCommend
	keyjson.Fun = "key.accept"
	keyjson.Client = "wheel"
	keyjson.Match = minionkey

	keyjson2, err := json.Marshal(keyjson)
	if err != nil {
		fmt.Println(err)
	}
	saltapi := beego.AppConfig.String("saltapi")
	requestjsoninfo, err2 := http.NewRequest("POST", saltapi, bytes.NewReader(keyjson2))
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
	if err4 != nil {
		fmt.Println(err4)
	}
	var jiange KeyAcceptedST

	err5 := json.Unmarshal([]byte(body), &jiange)

	if err5 != nil {
		fmt.Println(err5)

	}



	err98 := errors.New("无法添加此服务器")

	for _, v := range jiange.Return{
		a := v.Data.Return.Minions
		fmt.Println(a)

		fmt.Println(jiange)
		if s := len(a); s != 0 {

			fmt.Println("shujuwei......2222...",a)
			fmt.Println()
			return a, nil
			
		}
		return nil, err98
	}


	fmt.Println("shujuwei......3333...")
	return nil, nil

}

func KeyDeletAction(minion string)  bool {

	var keyjson  ActionCommend
	keyjson.Fun = "key.delete"
	keyjson.Client = "wheel"
	keyjson.Match = minion

	keyjson2, err := json.Marshal(keyjson)
	if err != nil {
		fmt.Println(err)
	}
	saltapi := beego.AppConfig.String("saltapi")
	requestjsoninfo, err2 := http.NewRequest("POST", saltapi, bytes.NewReader(keyjson2))
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
	if err4 != nil {
		fmt.Println(err4)
	}
	var jiange KeyDeletST

	err5 := json.Unmarshal([]byte(body), &jiange)

	if err5 != nil {
		fmt.Println(err5)

	}
	for _, v := range jiange.Return{

		a := v.Data.Success
		if a {
			return true
		}
		return false

	}
	return false

}









