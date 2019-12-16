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

1.结构体可以 使用嵌套匿名结构体所有的字段和方法，即：首字母大写或者小写的字段、方法，都可以使用。
2. 当 结构体和 匿名结构体有相同的字段或者方法时， 编译器采用就近访问原则访问，如希望访问
匿名结构体的字段和方法，可以通过匿名结构体名来区分
3. 结构体嵌入两个(或多个)匿名结构体，如 两个匿名结构体有相同的字段和方法( 同时结构体本身
没有同名的字段和方法)，在访问时，就必须明确指定匿名结构体名字，否则编译报错。
4. 如果一个 struct 嵌套了一个有名结构体，这种模式就是 组合，如果是组合关系，那么在访问组合
的结构体的字段或方法时，必须带上结构体的名字
6.  嵌套匿名结构体后，也可以在创建结构体变量(实例)时，直接指定各个 匿名结构体字段的值
*/




// todo 多重继承
/*
1. 如 一个 struct 嵌套了多个匿名结构体，那么该结构体可以直接访问嵌套的匿名结构体的字段和方
法， 从而实现了多重继承
如嵌入的匿名结构体有相同的字段名或者方法名，则在访问时，需要通过匿名结构体类型名来
区分
*/
type Goods struct {
	Name string
	Price float32
}
type Brand struct{
	Name string
	Address string
}
type Tv struct{
	//Goods
	//Brand
	a Goods
	b Brand
}