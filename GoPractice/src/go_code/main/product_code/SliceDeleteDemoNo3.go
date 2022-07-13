package main

import "fmt"

func main() {
	/*声明一个切片*/
	var array = []int{1, 2, 3, 4, 5, 6}

	/*使用copy函数进行复制*/
	array = array[:copy(array, array[1:])]
	array = array[:copy(array, array[3:])]

	fmt.Println(array)
	/*
	1、利用切片的性质，从开头获取元素
	2、使用copy函数返回一个int类型的数
		1、将要删除的开头元素不包含再目标切片当中
	 */
}