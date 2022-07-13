package main

import "fmt"

/*定义一个结构体,里面存储私有属性*/
type Circle struct {
	radius float64
}

//定义一个方法--->该 method 属于 Circle 类型对象中的方法
func (c Circle) getArea() float64 {
	return 3.14*c.radius*c.radius
}

func main() {
	/*声明一个结构体类型的变量，访问结构体的属性*/
	var c1 Circle
	c1.radius = 10.00

	/*因为getArea方法属于结构体，所以只能使用结构体变量去调用它*/
	fmt.Println("面积为:", c1.getArea())
}