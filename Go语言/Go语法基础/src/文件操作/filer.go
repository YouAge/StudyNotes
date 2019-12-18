package main

import (
	"fmt"
	"os"
)

// 文件操作
func openFile(name string)  {
	// 打开文件
	file,err := os.Open(name)
	defer file.Close()

	if err!= nil{
		fmt.Println("打开文件出错：",err)
	}else {
		fmt.Printf("file=%v",file)
	}
}

func main()  {

	path := "C:/Users/Master/Desktop/updater/extensions.py"
	openFile(path)

}

