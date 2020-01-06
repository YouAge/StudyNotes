package pubilc

import "fmt"

func IsError(err error)  {
	if err != nil{
		fmt.Println(err)
		return
	}


}
