package main

import (
	"fmt"
	"os"
	"chatRoom/pubilc"
)

// 定义2个全局变量
var userId int
var password string

func loging()  {
	fmt.Println("请输入用户id：")
	fmt.Scanf("%d\n",&userId)
	fmt.Println("请输入用户密码：")
	fmt.Scanf("%v\n",&password)
	err := login(userId,password)
	pubilc.IsError(err)

}

func main() {

	// 接收用户的选择
	var key int
	var loop = true

	for {
		fmt.Println("==========聊天室菜单==========")
		fmt.Println("\t\t\t\t1.登入聊天")
		fmt.Println("\t\t\t\t2.注册用户")
		fmt.Println("\t\t\t\t3.退出系统")
		fmt.Println("\t\t\t\t请选择（1-3）：")

		fmt.Scanf("%d\n", &key)

		switch key {
		case 1:
			fmt.Println("登入界面")
			loop = false

		case 2:
			fmt.Println("注册用户")
			loop = false
		case 3:
			fmt.Println("退出系统")
			//loop =false
			os.Exit(0)
		default:
			fmt.Println("输入有误，请重新选择")

		}
	}

}
