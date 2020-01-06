package main

import (
	"fmt"
	"time"
)

func exception(name string) {
	if err := recover(); err != nil {
		fmt.Printf("%v()函数线程出现错误", name)
	}

}

func sayHello()  {
	for i:=0;i<10;i++{
		time.Sleep(time.Second)
		fmt.Println("正常的协程")
	}
}

func test() {

	defer exception("test")
	var myMap map[int]string
	myMap[0] = "gloang"

}




func main() {
	//  默认情况下， 管道时双向的  （可读可写）
	// var chan1 chan int

	go sayHello()
	go test()

	// 声明只写
	var chan2 chan<- int
	chan2 = make(chan int, 3)
	chan2 <- 2
	//num := <-chan2    // cuow
	fmt.Println(chan2)

	// 声明为只读
	//var chan3 <-chan int
	//chan3 = make(chan int, 3)
	//num2 := <-chan3
	//fmt.Println(num2)

	//chan3 <- 30  写不进去

	// todo 使用select 可以解决从管道取数据阻塞问题
	selectG()
	time.Sleep(time.Second * 10)

}

func selectG() {
	// 定义管道
	intChan := make(chan int, 10)
	for i := 0; i < 10; i++ {
		intChan <- i
	}

	// 定义一个管道 5个数据sting
	strChan := make(chan string, 5)
	for i := 0; i < 5; i++ {
		strChan <- "hello" + fmt.Sprintf("%d", i)
	}

	// 在实际开发中，不确定 管道何时关闭，可以使用select
label:
	for {
		select {
		case v := <-intChan:
			fmt.Println("从intChan读取数据", v)

		case v := <-strChan:
			fmt.Println("从intChan读取数据", v)
		default:
			fmt.Println("没有数据了")
			break label
		}
	}

}
