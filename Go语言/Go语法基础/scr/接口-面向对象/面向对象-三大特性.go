package main

import "fmt"

// go 的OOP， 有所不同， 但是大体差不多

// TODO  封装
/*
封装(encapsulation)就是把抽象出的字段和对字段的操作封装在一起,数据被保护在内部,程序的其
它包只有通过被授权的操作(方法),才能对字段进行操作


// 封装的好处
// 1. 隐藏实现细节
// 2. 提可以对 数据进行验证，保证安全合理(Age)

如何体现封装：
1. 对结构体中的属性进行封装
2. 通过方法，包 实现封装


封装的实现步骤：

1) 将结构体、字段(属性)的首字母小写(不能导出了，其它包不能使用，类似 private)
2) 给结构体所在包提供一个工厂模式的函数，首字母大写。类似一个构造函数
3) 提供一个首字母大写的 Set 方法(类似其它语言的 public)，用于对属性判断并赋值
	func (var 结构体类型名) SetXxx(参数列表) (返回值列表) {
	//加入数据验证的业务逻辑
	var.字段 = 参数
	}
4) 提供一个首字母大写的 Get 方法(类似其它语言的 public)，用于获取属性的值
	func (var 结构体类型名) GetXxx() {
	return var.age;
	}
特别说明：在 Golang 开发中并没有特别强调封装，这点并不像 Java. 所以提醒学过 java 的朋友，
不用总是用 java 的语法特性来看待 Golang, Golang 本身对面向对象的特性做了简化的


外包不能直接反问首字母小写的函数，方法，变量等...
类似Python的私有属性...
*/

type person struct {
	Name string
	age  int
	sal  float64
}

// 工程模式函数， （其实就是调用桥梁）
func NewPerson(name string) *person {
	return &person{
		Name: name,
	}
}

//为了反问age和sal
func (p *person) SetAge(age int) {
	if age > 0 && age < 150 {
		p.age = age
	} else {
		fmt.Println("年龄不正确")
	}

}

func (p *person) GetAge() int {
	return p.age
}

func (p *person) SetSal(sal int) {
	if age > 3000 && age <= 30000 {
		p.age = sal
	} else {
		fmt.Println("薪水范围不正确")
	}

}

func (p *person) GetSal() float64 {
	return p.sal
}

// TODO 课堂练习：
/*
1) 创建程序,在 model 包中定义 Account 结构体：在 main 函数中体会 Golang 的封装性。
2) Account 结构体要求具有字段：账号（长度在 6-10 之间）、余额(必须>20)、密码（必须是六
3) 通过 SetXxx 的方法给 Account 的字段赋值。
4) 在 main 函数中测试

*/

type account struct {
	accountNa string
	pwd  int
	money    float64
}

func NewAccount(accountNa string,pwd int,money float64 ) *account  {
	if 6> len(accountNa) || len(accountNa)>10 {
		fmt.Println("输入账号格式不对")
		return nil
	}
	if pwd <6{
		fmt.Println("输入密码格式不对")
		return nil
	}
	if money <20{
		fmt.Println("金额最少20")
		return nil
	}
	return &account{
		accountNa: accountNa,
		pwd:       pwd,
		money:     money,
	}

}