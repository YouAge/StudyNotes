package main

import "fmt"

func TypeJude(items... interface{})  {
	for index,x :=range items{
		switch x.(type) {
		case bool:
			fmt.Printf("第%v个参数是 bool 类型，值是%v",index,x)
			index++
		case float32:
			fmt.Printf("第%v个参数是 float32 类型，值是%v\n",index,x)
			index++
		case float64:
			fmt.Printf("第%v个参数是 float64 类型，值是%v\n",index,x)
			index++
		case int,int32,int64:
			fmt.Printf("第%v个参数是 int,int32,int64 类型，值是%v\n",index,x)
			index++
		case string:
			fmt.Printf("第%v个参数是 string 类型，值是%v\n",index,x)
			index++
		default:
			fmt.Printf("第%v个参数是 不知道 类型，值是%v\n",index,x)
			index++
		}
	}
}

func main()  {
	var n1 float32=1.1
	var n3 float64=4.1
	var n2 int32 =30
	var name string  ="dsfs"
	addres:= "北京"
	n4 := 3000
	TypeJude(n1,n2,n3,name,addres,n4,"")



}