package main

import "fmt"

func main() {
	/*声明一个切片*/
	var array []int

	/*循环往里面赋值*/
	for i := 0; i < 10; i++ {
		/*使用append函数往切片里面添加元素*/
		array = append(array, i)
		fmt.Printf("len:%d,cap:%d,pointer:%p\n", len(array), cap(array), array)
		/*
		1、len函数查看切片拥有的元素个数
		2、cap函数查看切片容量情况
		 */
	}
}