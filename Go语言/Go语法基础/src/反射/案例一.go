package main

import (
	"fmt"
	"reflect"
)

// 使用反射，遍历结构体字段，调用结构体的方法，并获取结构体标签的值

// 定义一个Monster结构体
type Monster struct{
	Name string  `json:"name"`
	Age int `json:"age"`
	Score float32 `json:"成绩"`
	Sex string
}

// 返回2个数之和 方法
func (s Monster)GetSum(n1,n2 int) int {
	return n1+n2
}

func (s Monster) Set(name string,age int,score float32,sex string)  {
	// 给结构体赋值
	s.Name = name
	s.Age = age
	s.Score = score
	s.Sex = sex

}

func (s Monster) Print(){
	fmt.Println("---start---")
	fmt.Println(s)
	fmt.Println("---end---")
}

// 反射函数
func TestStruct(a interface{})  {
	// 获取reflect.Type类型
	typ := reflect.TypeOf(a)
	// 获取reflect.Value 类型
	val := reflect.ValueOf(a)
	// 获取 到 啊对应的类别
	kd := val.Kind()
	// 如果传入的不是struct，就退出

	if kd != reflect.Struct{
		fmt.Println("expect struct")
		return
	}





	// 获取到结构体的有几个字段
	num := val.NumField()

	fmt.Printf("struct has %d fields\n", num) //4

	// 变量结构体的所以字段
	for i :=0; i<num;i++{
		fmt.Printf("Field %d: 值为=%v\n", i, val.Field(i))
		// 获取struct 标签，注意需要通过 reflect.Type 来获取 tag 标签的值
		tagVal := typ.Field(i).Tag.Get("json")
		// 如果该字段于tag标签就显示， 否则就不显示
		if  tagVal !=""{
			fmt.Printf("Field %d: tag 为=%v\n", i, tagVal)
		}

	}

}

func main()  {

	a := Monster{
		Name:  "老鼠精",
		Age:   400,
		Score: 30.8,
		Sex:   "",
	}
	//TestStruct(a)
	TestStruct02(&a)

}

/*`
 使用反射的方式来获取结构体的 tag 标签, 遍历字段的值，修改字段值，调用结构体方法(要求：
通过传递地址的方式完成, 在前面案例上修改即可)
3) 定义了两个函数 test1 和 test2，定义一个适配器函数用作统一处理接口【了解】
4) 使用反射操作任意结构体类型：【了解】
5) 使用反射创建并操作结构体
`
*/

func TestStruct02(a interface{})  {
	b := a.(Monster)


	// 获取reflect.Type类型
	typ := reflect.TypeOf(b)
	// 获取reflect.Value 类型
	val := reflect.ValueOf(b)
	// 获取 到 啊对应的类别
	//kd := val.Kind()
	// 如果传入的不是struct，就退出

	//if kd != reflect.Struct{
	//	fmt.Println("expect struct")
	//	return
	//}


	//call := val.Interface().(*Monster) // 断言回结构体了吗？
	trueVar := val.Elem() // 获取真实的结构体
	trueVar.FieldByName("Age").SetInt(500)

	// 获取到结构体的有几个字段
	num := val.NumField()

	fmt.Printf("struct has %d fields\n", num) //4

	// 变量结构体的所以字段
	for i :=0; i<num;i++{

		fmt.Printf("Field %d: 值为=%v\n", i, val.Field(i))
		// 获取struct 标签，注意需要通过 reflect.Type 来获取 tag 标签的值
		tagVal := typ.Field(i).Tag.Get("json")
		// 如果该字段于tag标签就显示， 否则就不显示
		if  tagVal !=""{
			fmt.Printf("Field %d: tag 为=%v\n", i, tagVal)
		}

	}

}