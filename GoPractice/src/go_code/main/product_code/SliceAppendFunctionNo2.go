package main

import "fmt"

func main() {
	/*新建切片*/
	var array = []int{2,3,4}

	/*开头添加传参需要改变，并且需要对原数组进行压缩*/
	array = append([]int{0,1}, array...)
	/*
	这个操作再内存层面已经对原来的切片进行了复制扩容
	1、先复制原有的切片
	2、再创建一个更大的切片空间
	3、将添加的元素和原有的切片元素赋值进入到创建的切片空间
	 */

	fmt.Println(array)
}