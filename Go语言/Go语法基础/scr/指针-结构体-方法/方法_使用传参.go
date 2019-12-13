package main

import "fmt"

type Pster struct {
	//Name int
}


func (p Pster) tast(n1 ,n2 int) int {

	return n1+n2

}
//

type Circle struct {
	radius float64
}

// 值传递，效率比较低， 通常是 指针绑定
func (c Circle)area() float64 {
	return 3.14*c.radius*c.radius

}
// 提高效率
func (c *Circle)area2() float64 {
	fmt.Printf("c12:%p\n",c)
	c.radius =1
	return 3.14*c.radius*c.radius

}



func main()  {

	var p  Pster
	n1:= 10
	n2:=20
	res:=p.tast(n1,n2)
	fmt.Println(res)



	// 创建结构体
	//var c  Circle
	//fmt.Printf("c:%p\n",&c)
	c :=Circle{8.2}
	//fmt.Printf("c:%p\n",&c)
	//res2 :=c.area2()  // --> 正常是：res2:=(&c).area2() // go 底层会自动判读，如果是指针会自动处理， 底层优化了

	fmt.Println(&c)
}