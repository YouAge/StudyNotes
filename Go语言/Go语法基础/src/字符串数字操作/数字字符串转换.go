package main

import (
	"fmt"
	"strconv"
)

func main() {
	a := 100
	b := "100"
	var c int64 = 103
	v :=3.144564543132
	var v5 float32 =4.23245
	var y  float64

	// string int
	b1,err := strconv.Atoi(b)
	fmt.Println(b1,err)
	// string int64
	b2,err := strconv.ParseInt(b,10,64)
	fmt.Println(b2,err)

	// int string
	a1 := strconv.Itoa(a)
	fmt.Println(a1)

	// int64 string
	c1:=strconv.FormatInt(c,10)
	fmt.Println(c1)

	//float转string //
	v1 := strconv.FormatFloat(v,'E',-1,32) // float32
	v2 := strconv.FormatFloat(v,'e',-1,64) // float32
	fmt.Println(v1,v2)
	// float 截取小数点位数,
	y,_ = strconv.ParseFloat(fmt.Sprintf("%.2f",v5),64)
	fmt.Printf("y:%T,%T",y,v5)


}
