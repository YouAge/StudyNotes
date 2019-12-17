package main

import (
	"fmt"
	"time"
)

// go 语言 oop
// 抽象
/*
定义一个结构体时，实际上就是把一类事务的共有属性（字段）和行为（方法）提取出来 形成一个物理模型，
这种研究问题的方法称为抽象
*/

type Account struct {
	AccountNo string
	Pwd       string
	Balance   float64
}

// 方法

func (account *Account) Deposite(money float64, pwd string) {
	// 对比输入密码
	if pwd != account.Pwd {
		fmt.Println("输入密码错误")
		return
	} else {
		if money <= 0 {
			fmt.Println("请输入正确金额")
		} else {
			account.Balance += money
			fmt.Println("存款成功：", account.Balance)
		}
	}

}

//取款
func (account *Account) WithDraw(money float64, pwd string) {
	// 对比输入密码
	if pwd != account.Pwd {
		fmt.Println("输入密码错误")
		return
	} else {
		if money <= 0 || money > account.Balance {
			fmt.Println("请输入正确金额")
		} else {
			account.Balance -= money
			fmt.Println("取款成功：", account.Balance)
		}
	}

}

func (account *Account) Query(pwd string) {
	// 对比输入密码
	if pwd != account.Pwd {
		fmt.Println("输入密码错误")
		return
	} else {

		fmt.Println("余额为：", account.Balance)
	}
}

func main() {

	account := Account{
		AccountNo: "zhaoshang",
		Pwd:       "123456",
		Balance:   200.00,
	}

	for   {

		fmt.Println("菜单选项：\n" +
			"1.取款\n" +
			"2.存钱\n" +
			"3.查询")
		//mun:=

		account.Query("12346")
		account.Deposite(200,"123456")

		time.Sleep(2)
	}


}
