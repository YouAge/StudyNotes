package main

import (
	"chatRoom/pubilc"
	"fmt"
	"net"
)

const (

	NetWORk = "tcp"
	ip ="0.0.0.0"
	port = 8889
)


func process(conn net.Conn)  {
	// 读取客服端发送的信息
	defer conn.Close()

	for {
		buf:=make([]byte,8096)
		n,err := conn.Read(buf[:4])
		if n != 4 || err !=nil{
			fmt.Println("")
			return
		}
		fmt.Println("")
	}

}


func main()  {

	fmt.Println("服务器在")
	addr := fmt.Sprintf("%v:%d",ip,port)
	listen,err := net.Listen(NetWORk,addr)
	pubilc.IsError(err)
	defer listen.Close()
	for {
		fmt.Println("等待客服端连接")
		conn,err := listen.Accept()
		pubilc.IsError(err)
		go process(conn)
	}

}
