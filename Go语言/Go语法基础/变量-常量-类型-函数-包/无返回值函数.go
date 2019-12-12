package main

import "fmt"

// 无参数，返回值定义
func MyFunc() {
	a := 666
	fmt.Println("a=", a)
}
// 参数函数
func MyFun01(a ...int)  {
	//a = 111
	fmt.Println("a=",a[2])
}

func sum(n1 ,n2 int) int  {
	// 当执行到defer时，暂时不执行，会将defer后面的语句压入到独立的栈（defer栈）
	// 函数执行完， 再从defer栈 安 先入后出的方式取出，执行
	// 在defer将语句放入到栈时， 也会将相关的值深拷贝，入栈，
	defer fmt.Println("ok n1=",n1)  //3 -->10   // 拷贝了n1的值入栈，但当时是10,不会改变
	defer fmt.Println("ok2 n2=",n2) //2 -->20
	n1++
	n2++
	res := n1 +n2 //-->32
	fmt.Println("ok3 res=",res) // 1
	return res
}

func test()  {
	// filer = openfile("")
	// defer filer.close()
}




func main() {
	MyFun01(666,23452,23523)
	res := sum(10,20)
	fmt.Println("---",res)
	// 函数 defer  （在函数执行完后及时释放资源




}
