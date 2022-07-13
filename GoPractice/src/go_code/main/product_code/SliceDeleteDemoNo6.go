package main

import "fmt"

func main() {
	/*声明切片*/
	var array = []int{1, 2, 3, 4, 5, 6}

	/*对切片进行尾部删除--->使用len方法然后重新截取切片*/
	array = array[:len(array)-1]
	/*
	要删除第N个元素就-N
	 */

	fmt.Println(array)
}