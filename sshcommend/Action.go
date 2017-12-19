package sshcommend

import (
	"golang.org/x/crypto/ssh"
	"time"
	"fmt"
	"bytes"
	"net"
)


func Action(ip, user, pwd, commend string) (string, error) {

	var (
		client *ssh.Client
		session *ssh.Session
		err error
		a bytes.Buffer
	)
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
	if client, err = ssh.Dial("tcp",ip+":22", conf); err != nil{
		fmt.Println("zheliiiiiiiiiiiiiiiiii",err)
		return "", err
	}

	if session, err = client.NewSession(); err != nil{
		fmt.Println(err)
	}
	defer session.Close()

	session.Stdout = &a
	session.Run(commend)


	content := a.String()
	fmt.Println(content)
	return content, err
}

