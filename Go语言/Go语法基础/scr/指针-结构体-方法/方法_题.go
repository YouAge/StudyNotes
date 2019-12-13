package main

import "fmt"

//1) 编写结构体(MethodUtils)，编程一个方法，方法不需要参数，在方法中打印一个 10*8 的矩形，
//在 main 方法中调用该方法。
type MethodUtils struct {
}

func (ms1 *MethodUtils) ju() {
	for i := 0; i < 8; i++ {
		for j := 0; j < 10; j++ {
			fmt.Printf("*")
		}
		fmt.Println()
	}

}

//2. 编写一个方法，提供 m 和 n 两个参数，方法中打印一个 m*n 的矩形
func (ms1 *MethodUtils) area(m, n int) {
	fmt.Println("------------二-----------")
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			fmt.Printf("*")
		}
		fmt.Println()
	}

}

//3.编写一个方法算该矩形的面积(可以接收长 len，和宽 width)， 将其作为方法返回值。在 main
//方法中调用该方法，接收返回的面积值并打印
func (ms1 *MethodUtils) area2(m, n int) int {
	fmt.Println("------------三-----------")
	return m * n

}

//4. 编写方法：判断一个数是奇数还是偶数

func (ms1 *MethodUtils) oddEven(

	m int) {
	fmt.Println("------------四-----------")
	if m != 0 {
		if m%2 == 0 {

			fmt.Println("是偶数")
		} else {
			fmt.Println("是奇数")
		}
	} else {
		fmt.Println("0 不属于")
	}
}

//5. 根据行、列、字符打印 对应行数和列数的字符，比如：行：3，列：2，字符*,则打印相应的效 果
func (ms1 *MethodUtils) odd(m, n int, str string) {
	fmt.Println("------------五-----------")
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			fmt.Printf(str)
		}
		fmt.Println()
	}

}

/* 6.
定义小小计算器结构体(Calcuator)，实现加减乘除四个功能
实现形式 1：分四个方法完成:
实现形式 2：用一个方法搞定i
*/

func main() {
	// 1
	var p MethodUtils
	p.ju()
	//2

	//p2:=MethodUtils{}
	p.area(4, 5)

	//3
	//p3:=MethodUtils{}
	ftArea := p.area2(3, 4)
	fmt.Println(ftArea)

	// 4
	p.oddEven(7)

	//5
	p.odd(3, 5, "&&")

	var p1 *MethodUtils =&MethodUtils{}
	p1.odd(2,3,"0")


}
