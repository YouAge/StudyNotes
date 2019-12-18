package view

import (
	"fmt"
	"strconv"
	"接口-面向对象/客服管理系统/model"
	"接口-面向对象/客服管理系统/service"
)

type customerView struct {
	key             string
	loop            bool
	customerService *service.CustomerService
}

// 工程模式， 初始化
func NewCustomerVies() *customerView {
	return &customerView{
		key:             "",
		loop:            true,
		customerService: service.NewCustomerService(),
	}

}

func (vi *customerView) MainMenu() {
	for {
		fmt.Println("-----------------客服信息关系系统------------------------")
		fmt.Println("					1.添加客服信息")
		fmt.Println("					2.修改客服信息")
		fmt.Println("					3.删除客服信息")
		fmt.Println("					4.查看客服信息")
		fmt.Println("					5.退出客服信息")

		fmt.Println("请输入菜单指令：")
		fmt.Scanln(&vi.key)
		switch vi.key {

		case "1":
			vi.add()
		case "2":
			fmt.Println("修改客服信息")
		case "3":
			vi.delete()
		case "4":
			vi.list()
		case "5":
			vi.loop = false
		default:
			fmt.Println("无效的输入")
		}
		if !vi.loop {
			fmt.Println("欢迎下次再来")
			break
		}
	}

}

// 显示所有客服信息
func (vi *customerView) list() {

	coutomers := vi.customerService.List()
	fmt.Println("-------------客服列表展示----------------")
	fmt.Println("编号\t姓名\t性别\t年龄\t电话\t邮箱")
	for i := 0; i < len(coutomers); i++ {
		fmt.Println(coutomers[i].GetInfo())
	}
	fmt.Println("-------------客服列表完成----------------")
}

func (vi *customerView) add() {

	fmt.Println("---------添加客服-------------")
	fmt.Println("姓名：")
	name := ""
	fmt.Scanln(&name)
	fmt.Println("性别：")
	gender := ""
	fmt.Scanln(&gender)
	fmt.Println("年龄：")
	age := 0
	fmt.Scanln(&age)
	fmt.Println("电话：")
	phone := ""
	fmt.Scanln(&phone)

	fmt.Println("邮箱：")
	email := ""
	fmt.Scanln(&email)

	customer := *model.NewCustomer2(age, name, gender, phone, email)
	if vi.customerService.Add(customer) {
		fmt.Println("--添加成功--")
	} else {
		fmt.Println("--添加失败--")
	}

}

func (vi *customerView) delete() {
	// 删除指定信息
	for {
		index := ""
		fmt.Println("输入要删除人的编号，输入【m】查看列表，或者[n]返回菜单")
		fmt.Scanln(&index)
		if index == "n" {
			return
		}
		if index =="m"{
			// 查看
			vi.list()
			continue
		}
		id, err := strconv.Atoi(index)
		//fmt.Printf("err:%v,type:%T",err,err)
		if err == nil {
			if vi.customerService.Delete(id) {
				fmt.Println("-------删除成功")
				return
			} else {
				fmt.Println("--------删除失败，不存在此人，请重新操作")
			}
		} else {
			fmt.Println("请输入正确的序号---")
		}
	}
}
