package service

import (
	"接口-面向对象/客服管理系统/model"
)

type CustomerService struct {
	customers   []model.Customer // 结构体切片
	customerNum int              // 标号
}

//工程接口
func NewCustomerService() *CustomerService {
	// 为了能够看到客服信息， 初始化一个
	customerService := &CustomerService{}
	customerService.customerNum = 1
	customer := *model.NewCustomer(customerService.customerNum , 20, "张三", "男", "123141", "124@qq.com")
	//customer1 := *model.NewCustomer(1, 20, "小王", "男", "123141", "124@qq.com")
	customerService.customers = append(customerService.customers, customer)
	return customerService
}

// 返回客服切包， 就是全部是客服信息
func (se *CustomerService) List() []model.Customer   {
	return se.customers
}


// 添加 客服到customer 切片中
func (se *CustomerService) Add(customer model.Customer) bool {
	se.customerNum ++
	customer.Id = se.customerNum
	se.customers = append(se.customers,customer)
	return true
}


func (se *CustomerService) Delete(id int) bool {
	index := se.FindById(id)
	if index == -1{
		return false
	}
	// 切片中删除元素， 只能通过切片切割再合并
	se.customers =append(se.customers[:index],se.customers[index+1:]...)
	return true
}

// 跟去id查找客服，切片中对应的下标，没有返回-1
func (se *CustomerService) FindById(id int) int {

	index := -1
	for i:=0; i<len(se.customers);i++{
		if se.customers[i].Id == id {
			index =i
		}
	}
	return index
}