package main

import "fmt"

func main() {
	/*声明一个切片*/
	var array = []int{2,3,4,5,6}

	/*通过添加append的切片实现切片的拼接*/
	array = append(array[:3], append([]int{0,1}, array[1:]...)...)
	/*
	左闭右开：
	1、左边闭合的区间，左边的索引的值会取到该索引所在的位置
	2、右边开放的区间，右边索引的值会取到该索引前一位的位置
	 */

	fmt.Println(array)
	fmt.Println(array[:3])
	fmt.Printf("len:%d,cap:%d,pointer:%p", len(array), cap(array), array)
}