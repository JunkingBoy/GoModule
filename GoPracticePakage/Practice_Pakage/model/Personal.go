package model

import "errors"

/* 新建一个员工数据结构体 */
type person struct {
    Name string
    age int // 其他包不可直接访问
    salary float64
}

/* 工厂模式函数，相当于构造函数，返回结构的指针对象 */
func NewPerson(name string) *person {
    return &person{
        Name: name,
    }
}

/* 给person结构体的属性提供get和set方法 */
func (p *person) SetAge(age int) error {
    if age > 0 && age < 60 {
        p.age = age
    }else {
        return errors.New("年龄有误!")
    }
    return nil
}

func (p *person) GetAge() int {
    return p.age
}

func (p *person) SetSalary(sal float64) error {
    if sal > 3000 {
        p.salary = sal
    }else {
        return errors.New("薪资有误!")
    }
    return nil
}

func (p *person) GetSalary() float64 {
    return p.salary
}
