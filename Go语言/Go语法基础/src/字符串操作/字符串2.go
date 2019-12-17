package main

import (
	"fmt"
	"strconv"
	"strings"
)

func Start()  {
	fmt.Println("sdfsd")

}
/// 常用的字符串处理函数
//  内置函数  ---》 bulltin
func len_Str()  {
	// 统计字符串的长度，安字节 len（str）
	// golang的编码统一为utf-8 （ ascii的字符(字母和数字)个站一个，中文站3个字节）
	str :="lover"
	str1:="123中国"
	fmt.Println(len(str),"---",len(str1)) // 5,9

	str2 := "hello北京"
	r := []rune(str2) // 切片 ，分隔有中文的字符串
	for i:=0;i<len(r);i++{
		fmt.Printf("字符=%c\n",r[i])
	}

	// 字符串转整数，n,err:=strconv.Atoi("12")
	// 整数装字符串  str :=strconv.Itoa(12)
	// 字符串转byter   var byters =[]byte(""hello go)
	// byte转字符串 str = string([]byte(97,98,99))
	// 10 进制转，2,8,16  str = string.FormatInt(121,2)
	n,err := strconv.Atoi("12")
	if err !=nil{
		fmt.Println("失败：",n)
	}else {
		fmt.Printf("int=%T",n)
		fmt.Println("成功",n)
	}

	// 查找字串是否在制定的字符串中 -- 返回 true false
	b := strings.Contains("content","c")
	fmt.Println(b)
	// 统计字符串中有几个自定的子串
	num := strings.Count("cessese","s")
	fmt.Println(num)
	// 不区分大小写的字符串比较（==是区分大小写）（string.EqualFold("abc","ABC")）
	b1 := strings.EqualFold("abc","ABC")
	fmt.Println(b1) //true
	fmt.Println("结果：","abc"=="ABC") // flase

	// 返回子串第一次出现的index值，如果没有返回-1
	// strings.Index("Nlt_abc","abc") //4
	//  strings.LastIndex("Nlt_abc","abc")
	index :=strings.Index("Nlt_abc","abc")
	fmt.Println(index)

	// 将指定的字符替换： strings.Replact("go go hell","go","go语言"，n)
	// n 可以指定你希望替换几个，如果n=-1 全部替换
	str3 := "go go hell0"
	str4 := strings.Replace(str3,"go","hello",1)
	fmt.Println(str4)
	// 按照指定的某个字符，为分割标识，将一个字符拆分成字符串数组
	// strings.Split("hello,word,ok",",")
	strArr:=strings.Split("hello,word,ok",",")
	for i:=0;i<len(strArr);i++{
		fmt.Println(strArr[i])
	}
	// 将字符串的字母进行大小写装换 strings.ToLower("Go") //go
	// strings.ToUpper("gO") //GO

	//
}






func main()  {
	keys:= ""
	fmt.Println("1234")
	// len()函数
	len_Str()
	fmt.Println("输入：")
	key,err := fmt.Scanln(&keys)
	fmt.Println(key,err)
}