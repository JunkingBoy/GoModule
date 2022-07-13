package main

import "fmt"

/*
定义一个内嵌结构体
 */
type inner struct {
	num1 int
	num2 int
}

/*
定义另一个外部结构体嵌入内嵌结构体
 */
type outer struct {
	num3 int
	num4 float64
	//匿名结构体
	int
	//内嵌结构体
	inner
}

func main() {
	//实例化外部结构体
	outer1 := new(outer)
	//访问变量并且选择性的初始化
	outer1.num3 = 30
	outer1.int = 80
	outer1.num1 = 20 //这个成员在inner当中
	outer1.num2 = 30 //这个成员在inner当中
	//打印结果
	fmt.Printf("outer.num3 is:%d\n", outer1.num3)
	fmt.Printf("outer.int is:%d\n", outer1.int)
	fmt.Printf("outer.num1 is:%d\n", outer1.num2)
	fmt.Printf("outer.num2 is:%d\n", outer1.num2)
	fmt.Println("############################")
	//使用结构体字面量赋值
	outer2 := outer{6, 3.14, 60, inner{5, 8}}
	fmt.Printf("outer2 is:", outer2)
}