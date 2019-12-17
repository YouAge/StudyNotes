package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main()  {
	// 随机生成5个数，并反转, Python的切包反转不适用

	// rand.Int() 随机
	var intArr [5]int
	rand.Seed(time.Now().UnixNano()) // 不知道seed（）值，不会重新生成
	for i:=0;i<len(intArr);i++{

		intArr[i] = rand.Intn(1000)
	}
	fmt.Println("前：",intArr)
	//方法一： 必须是切片累成，才能使用
	//p:=intArr[:5]
	//sort.Slice(p, func(i, j int) bool {
	//	return true
	//})
	//fmt.Println("后：",p)

	// 方法二： 冒泡算法
	//for i,j:=0,len(intArr)-1;i<j;i,j=i+1,j-1{
	//	intArr[i],intArr[j] = intArr[j],intArr[i]
	//}
	//fmt.Println("后2：",intArr)
	//myInts := []int{ 8, 6, 7, 5, 3, 0, 9 }
	// 方法三 递归： 这是为了好玩而不是生产。 它不仅速度慢，而且如果列表太大，它将溢出堆栈。 我刚刚测试了，它将反转100万个int的列表但崩溃了1000万个
	bo:=reverseInts(intArr)
	fmt.Println("后3：",bo)


}

func reverseInts(input []int) []int {
	if len(input) ==0{
		return input
	}
	return append(reverseInts(input[1:]),input[0])
}