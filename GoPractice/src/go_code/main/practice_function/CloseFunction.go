package main

import "fmt"

func main() {
	//定义一个字符串
	str := "Hello World"

	//在这里先打印一遍str与后面的str打印做对比
	fmt.Println(str)

	//创建一个匿名函数
	foo := func() {
		//匿名函数中访问外部函数变量
		str = "Hello JunkingBoy"
	}

	//调用匿名函数
	foo() //这里调用str的值会被修改

	//打印出str
	fmt.Println(str)
}