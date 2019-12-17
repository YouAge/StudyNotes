package main

import "fmt"

type Person struct {
	name string
	age int
	flg bool
	intereset []string

}

func initFunc(p *Person)  {
	p.name = "Nami"
	p.age = 19
	p.flg =true
	p.intereset = append(p.intereset,"shopping","music")
}
func initFunc2(p *Person)  {
	p.name = "Imer"
	p.age = 16
	p.flg =true
	p.intereset = append(p.intereset,"shopping","music")
}



func main()  {

 // 结构体常使用，指针传递，  值传递会生产大量内存，一般不用
	var  person Person
	initFunc(&person)
	initFunc2(&person)
	fmt.Println(person)

	//append(a,person)

	// 放入地址，
	num := new(int)  // 开辟空间， num是地址，*num取值
 	fmt.Println(*num,&num)
	*num =100
	fmt.Println(*num,&num)

}
