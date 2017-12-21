package sshcommend

import (
	"golang.org/x/crypto/ssh"
	"time"
	"fmt"
	"bytes"
	"net"
)

//ssh连接方法
func Action(ip, user, pwd, commend string) (string, error) {

	var (
		client *ssh.Client
		session *ssh.Session
		err error
		a bytes.Buffer
	)
	//初始化客户端连接参数，配置
	conf := &ssh.ClientConfig{
		User:user,
		Auth:[]ssh.AuthMethod{
			ssh.Password(pwd),
		},
		Timeout: 30 * time.Second,
		HostKeyCallback: func(hostname string, remote net.Addr, key ssh.PublicKey) error {
			return nil
		},
	}
	fmt.Println(ip+":22")
	//开始建立客户端tcp握手连接对象
	if client, err = ssh.Dial("tcp",ip+":22", conf); err != nil{
		fmt.Println("zheliiiiiiiiiiiiiiiiii",err)
		return "", err
	}
    //创建会话，为这个客户端，可以执行一个程序
	if session, err = client.NewSession(); err != nil{
		fmt.Println(err)
	}
	defer session.Close()
    //会话标准输出到a变量内
	session.Stdout = &a
	//执行命令
	session.Run(commend)

    //打印标准输出信息
	content := a.String()
	fmt.Println(content)
	return content, err
}

