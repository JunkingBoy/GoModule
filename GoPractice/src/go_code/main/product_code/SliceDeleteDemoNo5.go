package main

import "fmt"

func main() {
	/*声明一个切片*/
	var array = []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	/*使用copy的方法*/
	//删除3号元素--->先把3号之后的元素复制到3号的位置上
	copy(array[:2], array[3:])
	/*
	这个时候3号元素之前的元素已经被删除了
	要再用切片的方法去保留前两位元素
	 */

	fmt.Println(array)

	array = []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	array = array[:2+copy(array[:2], array[3:])]

	fmt.Println(array)
}