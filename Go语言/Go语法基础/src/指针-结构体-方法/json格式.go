package main

import (
	"encoding/json"
	"fmt"
)

type Monster struct {
	Name     string `json:"name"`  // 指定标签，
	Age      int `json:"age"`
	Birthday string `json:"birthday"`
	Sal      float64 `json:"sal"`
	Skill    string
}


type Monster2 struct {
	Name     string `json:"name"`  // 指定标签，
	Age      int `json:"age"`
	Birthday string `json:"birthday"`
	Sal      float64

}

func testStruct() {
	monster := Monster{
		Name:     "孙悟空" ,
		Age:      6000,
		Birthday: "2011-12-03",
		Sal:      3444.66,
		Skill:    "筋斗云",
	}
	// 序列化,  结构体，首字母小写字母 不会被序列号
	data,err :=json.Marshal(&monster)
	if err!=nil{
		fmt.Println("错误",err)
	}
	fmt.Println(string(data))

	// 反序列化
	str := `{"name":"孙悟空","age":6000,"birthday":"2011-12-03","sal":3444.66,"Skill":"筋斗云"}`
	//var c Monster2    // 结构体
	var c map[string]interface{}  // map ， 反序列化不用make， Unmarshal里面已经实现了make
	err2:=json.Unmarshal([]byte(str),&c) // &c 引用传递 改变值
	if err !=nil{
		fmt.Println("err",err2)
	}
	fmt.Println(c)
}







// 反序列化



func main() {
	testStruct()

}
