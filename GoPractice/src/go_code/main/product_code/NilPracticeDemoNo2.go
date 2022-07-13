package main

import "fmt"

func main() {
	/*声明两个变量*/
	var arr []int
	var num *int
	var num2 *int

	/*分别利用格式化输出他们的nil值*/
	fmt.Printf("%p\n", &arr)
	fmt.Printf("%p\n", &num)
	fmt.Printf("%p", &num2)
	/*
	fmt.Printf方法返回一个整数类型的值
	和常规的空值验证
	 */
}