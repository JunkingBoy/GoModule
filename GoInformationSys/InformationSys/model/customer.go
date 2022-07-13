package model

import "fmt"

/*
1、定义customer结构体
2、提供工厂构建结构体实例的函数
3、提供类似java的toString方法，获取结构体信息
*/
type Customer struct {
    Id int
    Name string
    Sex string
    Age int
    Phone string
    Email string
}

// 提供工厂构建customer结构体的函数
func NewCustomerByAll(id int, name string, sex string, age int, phone string, email string) Customer {
    // 返回结构体信息
    return Customer{
        Id: id,
        Name: name,
        Sex: sex,
        Age: age,
        Phone: phone,
        Email: email,
    }
}

func NewCustomerByPart(name string, sex string, age int, phone string, email string) Customer {
    // 返回结构体
    return Customer{
        Name: name,
        Sex: sex,
        Age: age,
        Phone: phone,
        Email: email,
    }
}

// 提供获取结构体信息的函数
func (this Customer) GetInfo() string {
    info := fmt.Sprintf("%v\t\t %v\t %v\t\t %v\t\t %v\t %v\t", this.Id, this.Name, this.Sex, this.Age, this.Phone, this.Email)
    return info
}
