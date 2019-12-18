package model

import "fmt"

// 声明一个Customer 结构体， 客户信息

type Customer struct {
	Id     int
	Name   string
	Gender string
	Age    int
	Phone  string
	Email  string
}

func NewCustomer(id, age int, name, gender, phone, email string) *Customer {
	// 第一种，手动添加id
	return &Customer{
		Id:     id,
		Name:   name,
		Gender: gender,
		Age:    age,
		Phone:  phone,
		Email:  email,
	}
}

func NewCustomer2( age int, name, gender, phone, email string) *Customer {
	// 第二种，自动id
	return &Customer{
		Name:   name,
		Gender: gender,
		Age:    age,
		Phone:  phone,
		Email:  email,
	}
}







// 放回用户的信息， 格式化
func (md Customer) GetInfo() string {
	info := fmt.Sprintf("%v\t%v\t%v\t%v\t%v\t%v\t",
		md.Id, md.Name, md.Gender, md.Age, md.Phone, md.Email)
	//a = info
	return info
}
