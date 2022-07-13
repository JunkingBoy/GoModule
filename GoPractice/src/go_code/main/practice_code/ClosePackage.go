package main

import "fmt"

func getSequence() func() int {
	i := 0
	//返回值书写内置函数
	return func() int {
		i += 1
		return i
	}
	/*
	内部返回值是int类型
	返回后外部返回值认为函数相当于一个int类型的引用变量
	 */
}

func main() {
	/*设置函数变量*/
	nextNumber := getSequence()

	/*调用nextNumber函数*/
	fmt.Println(nextNumber())
	fmt.Println(nextNumber())
	fmt.Println(nextNumber())
}