package main

import "fmt"

func main() {
	//字符串型切片
	var strList []string
	
	//整型切片
	var numList []int

	//空切片--->空数组是对变量赋值不是声明变量
	var emptyList = []int{} //emptyList 已经被分配了内存，只是还没有元素。

	fmt.Println(strList, numList, emptyList)

	fmt.Println(len(strList), len(numList), len(emptyList))

	fmt.Println(strList == nil)
	fmt.Println(emptyList == nil)
	/*切片是动态结构，只能与 nil 判定相等，不能互相判定相等。声明新的切片后，可以使用 `append()` 函数向切片中添加元素*/
}