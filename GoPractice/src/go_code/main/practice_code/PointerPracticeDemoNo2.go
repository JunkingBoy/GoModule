package main

import "fmt"

func main() {
	var house = "home"

	//地址赋值指针
	ptr := &house

	//打印指针类型
	fmt.Printf("Type is:%T\n", ptr)
	//打印地址
	fmt.Printf("Address is:%p\n", ptr)

	//使用指针指向指针
	value := *ptr //这个ptr本身的类型是*string
	/*实际上value的类型不是**string。这样赋值以后是直接拿到了指针ptr指向的内存地址的值*/

	//打印类型和值
	fmt.Printf("Value type is:%T\n", value)
	fmt.Printf("Vlaue is:%s\n", value)

	//声明一个指针
	var ptr2 **string
	ptr2 = &ptr

	fmt.Printf("Type is:%T\n", ptr2)
	fmt.Printf("Value is:%d", ptr2)
}