package main

import "fmt"

func main() {
	/*声明两个数组切片*/
	slice1 := []int{1,2,3,4,5}
	slice2 := []int{7,8,9}

	/*使用copy函数对他们的元素进行操作*/
	copy(slice1, slice2)
	slice1[0] = 1
	slice1[1] = 2
	slice1[2] = 3
	copy(slice2, slice1) //这个已经改变了数组当中的元素--->copy函数

	fmt.Println(slice1)
	fmt.Println(slice2)
}