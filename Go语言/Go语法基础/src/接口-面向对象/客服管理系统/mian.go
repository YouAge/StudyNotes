package main

import "接口-面向对象/客服管理系统/view"







func main()  {
	// 项目启动
	customerView := view.NewCustomerVies()

	customerView.MainMenu()

}
