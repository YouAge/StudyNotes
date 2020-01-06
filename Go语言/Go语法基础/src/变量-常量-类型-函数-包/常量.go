package main

import "fmt"

// 常量声明时必须给值
// 常量声明后，不能修改
// 常量只能修饰bool int float string  基本类型
// 常量仍然通过首字母大写 反问权限
const stu = "ssdfds"

func main()  {

	const name =5
	// name = 4
	const jl bool = true
	const g = 9/3  //(定值，不可变的)
	fmt.Println(g)

	const  (
		a = iota // 0  每一行递增加一
		b = iota
		j =4
		c,d = iota,iota

	)
	fmt.Println(a,b,c,d)
}
