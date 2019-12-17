package main

import (
	"fmt"
	g "接口-面向对象/账单项目"
)

/*
todo  简介
Golang 中 多态
特性主要是通过接口来体现的

interface(接口类型) 类型可以定义一组方法，但是这些不需要实现。 todo  并且 interface 不能包含任何变量。到某个
自定义类型(比如结构体 Phone)要使用的时候,在根据具体情况把这些方法写出来(实现)
基本语法：
	type name interface{
		method1(参数列表)返回值列表
		method2(参数列表)返回值列表

	}
	func (t 自定义类型)method1(参数列表)返回值列表{

	}

1.接口里的 所有方法都没有方法体，即接口的方法都是没有实现的方法。接口体现了程序设计的
多态和 高内聚低偶合 的思想
2.Golang 中的接口， 不需要 显式的实现。只要一个变量，含有接口类型中的所有方法，那么这个
变量就实现这个接口。因此，Golang 中 没有 implement 这样的关键字
*/

/*
todo 接口注意事项
1. 接口本身 不能创建实例,但是 可以指向一个实现了该接口的自定义类型的变量(实例)
2. 接口中所有的方法都没有方法体,即都是没有实现的方法。
3. 在 Golang 中，一个自定义类型需要将某个接口的所有方法都实现，我们说这个自定义类型实现
了该接口。 (缺一不可)
4.只要是自定义数据类型，就可以实现接口，不仅仅是结构体类型

*/

type AInterface interface {
	Say()
}
type Stu struct {
	Name string
}

func (stu Stu) Say() {
	stu.Name = "sdfs"
	fmt.Println("Stu Say()", stu.Name)
}

//4
type integer int

func (i integer) Say() {
	fmt.Println("integer Say()", i)
}

func mainTwo() {
	var stu Stu            // 接口变量实现
	var a AInterface = stu //可以指向一个实现了该接口的自定义类型的变量(实例)
	a.Say()

	var i integer = 10
	var c AInterface = i

	c.Say()
}

func main() {
	mainTwo()
	// 项目，调包
	b :=g.NewFamilyAccount()
	b.MainMenu()
}
