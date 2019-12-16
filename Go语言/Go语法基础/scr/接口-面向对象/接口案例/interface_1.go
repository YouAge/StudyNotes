package main

import (
	"fmt"
	"math/rand"
	"sort"
)

/*

type Alnterface interface {
	Test01()
	Test02()

}
type Blnterface interface {
	Test01()
	Test03()
}
type Stu struct {
}

func (stu Stu)Test01()  {}
func (stu Stu)Test02()  {}
func (stu Stu)Test03()  {}
func main() {
	stu :=Stu{}
	var a Alnterface =stu
	var b Alnterface =stu
	fmt.Println(a,b)
}


*/


// todo 排序
//实现对hero结构体排序

type Hero struct {
	Name string
	Age int
}
// 定义 结构体切片
type HeroSlice []Hero

func (hs HeroSlice) Len() int  {
	return len(hs)
}
func (hs HeroSlice)Less(i,j int) bool  {
	// 用什么字段排序，
	return hs[i].Age<hs[j].Age
	
}
func (hs HeroSlice) Swap(i,j int)  {
	//temp := hs[i]
	//hs[i] = hs[j]
	//hs[j] = temp
	// ==> 等价
	hs[i],hs[j] = hs[j],hs[i]
}

func mainSlice()  {
	var hs HeroSlice
	for i:=0;i<10;i++{
		sp := Hero{
			Name:fmt.Sprintf("姓名：%d",rand.Intn(100)),
			Age:rand.Intn(100),
		}
		hs = append(hs,sp)
	}


}



func main()  {

	var intSlice = []int{14,21,3,0,8,4}
	sort.Ints(intSlice) // 排序
	fmt.Println(intSlice)

	// d结构体切片排序

}



