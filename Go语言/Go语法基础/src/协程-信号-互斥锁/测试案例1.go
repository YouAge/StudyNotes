package main

import "fmt"

var m map[string] chan bool

func Sear1(intChan chan int)  {

	for i:=0;i<1000 ;i++  {
		fmt.Println("线程一")
	}
	intChan <- 1
	close(intChan)
}

func Sear2(intChan chan int)  {

	for i:=0;i<1000 ;i++  {
		fmt.Println("---线程2")
	}
	intChan <- 2
	close(intChan)
}

func main()  {
	name := make(chan int,2)
	go Sear1(name)
	go Sear2(name)
	<- name

}