package main

import (
	"container/list"
	"fmt"
)

func main() {
	/*声明一个列表*/
	listPractice := list.New()

	fmt.Println(listPractice)

	/*在列表后面插入一个值*/
	listPractice.PushBack("first")
	listPractice.PushBack(66)

	fmt.Println(listPractice) //这个打印会打印处列表当中的值的地址
}