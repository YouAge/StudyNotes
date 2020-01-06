package main

import (
	"chatRoom/pubilc"
	"encoding/json"
	"fmt"
	"net"
)

const   (

	NetWORk = "tcp"
	ip ="localhost"
	port = 8889
)

func login(userId int,password string) (err error)  {

	// 定义协议
	fmt.Printf("userId=%d,password=%v",userId,password)

	radds := fmt.Sprintf("%v:%d",ip,port)
	conn,err :=net.Dial(NetWORk,radds)
	pubilc.IsError(err)
	defer conn.Close()

	var mes pubilc.Message
	mes.Type = pubilc.LoginMesType

	var loginMes pubilc.LoginMes
	loginMes.UserId = userId
	loginMes.Password= password

	data,err := json.Marshal(loginMes)
	pubilc.IsError(err)

}