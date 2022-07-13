package service

import (
	"GoInformationSys/InformationSys/model"
)

/*
提供对customer结构体的操作
增、删、改、查
*/

// 提供一个结构体，里面包含当前切片
type CustomerService struct {
	// 声明一个字段，表示当前切片含有多少个客户--->用户结构体切片
	customers []model.Customer

	// 用户id
	customerNum int
}

// 返回customerservice结构体指针，展示默认的数据
func NewCustomerService() *CustomerService {
	// 初始化一个默认的客户
	customerService := &CustomerService{}
	// 设置默认用户id
	customerService.customerNum = 1

	// 设置默认用户信息--->调用customer包下的new函数
	customer := model.NewCustomerByAll(1, "JunkingBoy", "Man", 22, "18154223156", "Lucifer@test.com")
	// 将用户信息放到customerService切片中
	customerService.customers = append(customerService.customers, customer)
	// 返回customerService结构体
	return customerService
}

//// 返回customer指针
//func (this *CustomerService) GetSelfInfo() *CustomerService {
//    return &CustomerService{}
//}
//
//// 设置编号函数
//func (this *CustomerService) SetNum(id int) {
//    this.GetSelfInfo().customerNum = id
//}

// 获取customer切片的函数
func (this CustomerService) List() []model.Customer {
	return this.customers
}

// 添加客户到customers切片中
func (this *CustomerService) Add(customer model.Customer) bool {
	// 重新设置排序--->避免以及存在被删除的数据在自增的话会跳过一位数
	/* 获取customer长度、拿到他的排序 */
	customers := this.List()
	// 获取长度
	getLengh := len(customers)
	for i := 0; i <= getLengh; i++ {
		if i == getLengh {
			this.customerNum = i
		}
	}
	// 首先customerService的id自增
	this.customerNum++
	// 将传入的customerService的id设置为customer的id
	customer.Id = this.customerNum
	// 将形参放到切片中
	this.customers = append(this.customers, customer)
	return true
}

// 根据id查找客户--->没查到返回-1
func (this *CustomerService) FindById(id int) int {
	index := -1
	if id > 0 {
		// 循环遍历切片
		for i := 0; i < len(this.customers); i++ {
			if this.customers[i].Id == id {
				// 有该客户，找到该客户，返回对应的id
				index = i
				// 重新排序num
				/* 获取customer长度、拿到他的排序 */
				customers := this.List()
				// 获取长度
				getLengh := len(customers)
				for i := 0; i <= getLengh; i++ {
					if i == getLengh {
						this.customerNum = i
					}
				}
			}
		}
	}
	return index
}

// 根据id删除客户
func (this *CustomerService) Delete(id int) bool {
	if 0 < id {
		targetId := this.FindById(id)
		if targetId != -1 {
			// 从切片中删除元素
			this.customers = append(this.customers[:targetId], this.customers[targetId+1:]...)
			return true
		}
	}
	return false
}
