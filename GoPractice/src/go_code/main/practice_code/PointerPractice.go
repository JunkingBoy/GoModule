package main

import "fmt"

func main() {
	/*声明两个变量*/
	var num int = 1
	str := "香蕉"

	fmt.Printf("%p,%p\n", &num, &str)
	/*
	这只是获取了变量的内存地址
	 */

	/*声明指针并且赋值*/
	var ptr *int

	ptr = &num

	/*打印ptr观察是否与&num值一致*/
	fmt.Printf("%p", ptr)
}