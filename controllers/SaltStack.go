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

func Exec_commend(zhuji string, commend string)  (CommendRS , error) {

	var action ActionCommend
	var jiange CommendRS
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
	tocken, err := Tokend()
	if err != nil {
		return jiange, err
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

		fmt.Println(err3)
	}
	body ,err4 := ioutil.ReadAll(resp.Body)


	jsonerr := json.Unmarshal([]byte(body), &jiange)
	if jsonerr != nil {
		fmt.Println(jsonerr)
	}
	if err4 != nil {
		fmt.Println(err4)
	}
	fmt.Println(jiange)

	return jiange, nil


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

	jieguo , err := Exec_commend(zhujixinxi, commend)
	if err != nil {

		c.Data["commedinfo"] = "此服务器异常"

		c.TplName = "saltremoteexecution.html"
		return

	}


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

type CpfileController struct {
	beego.Controller
}


type Cprespone struct {
	Return []map[string]string `json:"return"`
}


//curl -k https://133.130.122.48:8001/ -H "X-Auth-Token: 10f6ebdabb0c10730194eafbb071f05f2ade6638" -H "Accept: application/json" -d client=local -d tgt='163-44-167-92.conoha.io' -d fun='cp.get_file' -d arg='salt://hello2' -d arg='/tmp/hello2'

type CPrequest struct {
	Client string `json:"client"`
	Tgt string `json:"tgt"`
	Fun string `json:"fun"`
	Arg []string `json:"arg"`

}


func Cpfile(host, srfile, drpath string) (Cprespone ,error) {

	var action CPrequest
	action.Client = "local"
	action.Tgt = host
	action.Fun = "cp.get_file"
	action.Arg = append(action.Arg, "salt://" + srfile)
	action.Arg = append(action.Arg, drpath)
	actionjson, err := json.Marshal(action)
	fmt.Println("json info ......", string(actionjson))
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

		fmt.Println("cpfile tokenhuoqu shibai")
		//return nil, xinxi
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

		fmt.Println("cpfile baocuo ...", err3)
	}
	body ,err4 := ioutil.ReadAll(resp.Body)
	if err4 != nil {
		fmt.Println("cpfile read error........",err4)
	}
	fmt.Println(string(body))

	var jiange Cprespone

	jsonerr := json.Unmarshal([]byte(body), &jiange)

	if jsonerr != nil {
		fmt.Println(jsonerr)
	}

	return jiange, nil
	
	
	
}

func (c *CpfileController) UpDataFile() {


	f, h, err := c.GetFile("fileupdata")
	if err !=nil {
		fmt.Println("上传出错", err)
	}

	filename := h.Filename
	defer f.Close()
	c.SaveToFile("fileupdata", "D:\\bao\\" + filename)

	c.Data["Drfile"] = "ok"
	c.TplName = "saltfileCp.html"

}

func (c *CpfileController) CpfileGet() {


	c.TplName = "saltfileCp.html"
}

func (c *CpfileController) CpfileAction() {

	tgt := c.Input().Get("drhostname")
	sourcefile := c.Input().Get("srfile")
	tgtpath := c.Input().Get("drpath")


	response, err := Cpfile(tgt, sourcefile, tgtpath)

	if err !=nil {
		fmt.Println("脚本分发失败")
	}
	fmt.Println(response.Return)


	for _, v := range response.Return {

		var aa []string

		for k, v2 := range v{
			fmt.Println(k,v2)
			aa = append(aa, k+","+v2)

		}

		c.Data["Drfile"] = aa
		c.Data["DRhost"] = tgt
		c.Data["Sorfile"] = sourcefile
		c.Data["Drpath"] = tgtpath
	}

	c.TplName = "saltfileCp.html"

}






