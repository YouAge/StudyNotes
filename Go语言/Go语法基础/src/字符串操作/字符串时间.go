package main

import (
	"fmt"
	"strconv"
	"time"
)
//const (
//	Nanosecond  Duration = 1  // 纳秒
//	Microsecond          = 1000 * Nanosecon //微妙
//	Millisecond          = 1000 * Microsecond //毫秒
//	Second               = 1000 * Millisecond
//	Minute               = 60 * Second
//	Hour                 = 60 * Minute
//)

func testTime()  {
	str := ""
	for i :=0;i<1000000;i++{
		str +="hello"+strconv.Itoa(i)
	}
}



func main()  {

	fmt.Println("123")
	// 获取当前时间
	now := time.Now()
	fmt.Printf("now=%T and now=%v\n",now,now)

	// 2. 通过now 获取年月日，时分秒
	fmt.Printf("年=%v\n",now.Year())
	//fmt.Printf("月=%v\n",now.Month()) // December
	fmt.Printf("月=%v\n",int(now.Month()))

	fmt.Printf("日=%v\n",now.Day())
	fmt.Printf("时=%v\n",now.Hour())
	fmt.Printf("分=%v\n",now.Minute())
	fmt.Printf("秒=%v\n",now.Second())

	//格式化 日期时间
	fmt.Printf("当前时间： %d-%d-%d %d:%d:%d \n",now.Year(),now.Month(),
		now.Day(),now.Hour(),now.Minute(),now.Second(),
		)
	dateStr := fmt.Sprintf("当前时间： %d-%d-%d %d:%d:%d \n",now.Year(),now.Month(),
		now.Day(),now.Hour(),now.Minute(),now.Second(),
	)
	fmt.Println(dateStr)

	// 方法二
	fmt.Printf(now.Format("2006/01/02 15:04:05\n"))  // 数字不能更改，符号可以随意改， 不然格式不出
	fmt.Printf(now.Format("2006-01-02 15:04:05\n"))
	fmt.Printf(now.Format("2006-01-02\n"))
	fmt.Printf(now.Format("01\n")) //只拿月

	// 时间的常量，   在程序运行中 获取指定的时间
	i :=0
	for {
		i++
		fmt.Println(i)
		// 0.1秒 不能是Second*0.1不然不通过
		//time.Sleep(time.Second)
		time.Sleep(time.Microsecond * 100)
		if i==10{
			break
		}
	}
	// Unix 时间戳 秒的
	// UnixNano  纳秒时间戳
	fmt.Printf("unxs时间戳=%v，UnixNano时间戳=%v",now.Unix(),now.UnixNano())

	star:=time.Now().Unix()
	testTime()
	end:=time.Now().Unix()
	fmt.Printf("\n%v秒\n",end-star)


}