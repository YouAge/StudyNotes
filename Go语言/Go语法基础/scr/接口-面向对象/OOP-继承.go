package main

import "fmt"

// TODO 面向对象 -- 继承

/*
1. 代码冗余， 不利于维护，也不利于功能扩展，代码复用性不强

继承就是解决上面的问题

go 中的继承 就是 结构体 里面嵌套结构体
基本语法：
type Goods struct {
	Name string
	Price int
}
type Book struct {
	Goods //这里就是嵌套匿名结构体 Goods
	Writer string
}

*/

// 编写一个学生考试系统

type Student struct {
	Name  string
	age   int
	Score float32
}

// 将Pupil 和Graduate 共有的方法绑定到*Student

func (s *Student) ShowInfo() {
	fmt.Printf("学生：%v，年龄：%v , 正在考试", s.Name, s.age)
}
func (s *Student) SetScore(score float32) {
	s.Score =score
}

type Pupil struct {
	Student
}
//这时 Pupil 结构体特有的方法，保留
func (p *Pupil) testing() {
	fmt.Println("小学生正在考试中.....")
}

type Graduate struct {
	Student
}
//这时 Graduate 结构体特有的方法，保留
func (p *Graduate) testing() {
	fmt.Println("大学生正在考试中.....")
}



func main() {


	// 当 我们对结构体嵌入了匿名结构体使用方法会发生变化
	pupil := &Pupil{}
	pupil.Student.Name = "tom~"
	pupil.Student.age = 8
	pupil.testing()
	pupil.Student.SetScore(70)
	pupil.Student.ShowInfo()





}

/*
// todo 深入解析





*/
