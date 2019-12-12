package main

import (
	"fmt"
	"strings"
)

func main() {
	// 按 指定分隔符拆分


	str := "I love my work and I love my family too"

	// 按空格拆分
	ret := strings.Fields(str)
	//for _,key := range ret {
	//
	//}
	fmt.Println(ret)
	// 判断字符串字符串结束标记
	flg := strings.HasSuffix("sdfsdfs.abc","abc")
	fmt.Println(flg)

	//判断字符串起始标识
	str2 := `sdfsadfhisadgoahsoidghoiwahoigf
asdjgpofasjpodgjposdj	fmt.Println(ret)
	// 判断字符串字符串结束标记
	flg := strings.HasSuffix("sdfsdfs.abc","abc")
	fmt.Println(flg)gposdpogspojsdpojgspojgpoj"sfw"`
	fmt.Println(str2)


}
