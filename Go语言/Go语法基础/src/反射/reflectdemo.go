package main

import (
	"fmt"
	"reflect"
)

// 反射案例1
func reflectTest(b interface{})  {

	// t通过反射获取传入的变量， type kind 值
	// 1. 先获取 reflect.Type
	rTyp := reflect.TypeOf(b)
	fmt.Println("rTyp=",rTyp)

	// 获取reflect.Value
	rVal := reflect.ValueOf(b)

	// 获取对应kind的常量, 2种
	// rVal.Kind()
	// rTyp.kind()
	kind1 := rVal.Kind()
	kind2 := rTyp.Kind()
	fmt.Printf("kind=%v,kind=%v\n",kind1,kind2)


	n2 := 2+rVal.Int()
	fmt.Println(n2)

	fmt.Println("rVal=",rVal)
	fmt.Printf("rVal=%v，type=%T\n",rVal,rVal)


	// 将 rval 转成interface{},转回原类型
	iv := rVal.Interface()
	num2 := iv.(int) // 断言
	fmt.Println(num2)
}


func reflectTest02(b interface{}) {

// 对结构体反射
	// 1. 先获取 reflect.Type
	rTyp := reflect.TypeOf(b)
	fmt.Println("rTyp=",rTyp)

	// 获取reflect.Value
	rVal := reflect.ValueOf(b)



	// 将 rval 转成interface{},转回原类型
	iv := rVal.Interface()

	fmt.Printf("iv=%v,type=%T\n",iv,iv)
	// 将interface{} 通过断言转成需要的类型, 通过swtich 断言进行断言遍历判断

	stu,ok :=iv.(Student)
	if ok {
		fmt.Println(stu.Name,ok)
	}


}

type Student struct {
	Name string
	Age int
}


func main()  {

 //编写基本类型 进程反射

	num := 100
	reflectTest(num)
	reflect03(&num)
	fmt.Println(num)

	// 2. 定义一个结构体
	stu := Student{
		Name: "liyao",
		Age:  20,
	}

	reflectTest02(stu)
}



// 反射的注意事项，细节





// 通过反射修改值
func reflect03(b interface{})  {

	// 获取reflect.Value
	rVal := reflect.ValueOf(b)
	// 修改地址的值，rVal.Elem() == 》 *b
	rVal.Elem().SetInt(5) // 指针修改


}