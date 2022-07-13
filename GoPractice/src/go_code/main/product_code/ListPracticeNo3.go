package main

import (
	"container/list"
	"fmt"
)

func main() {
	/*声明一个列表*/
	listPractice := list.New()

	//尾部添加元素
	listPractice.PushBack("back")
	//头部添加元素
	listPractice.PushFront(66)

	//因为返回的类型是一个*element句柄，所以保存下来
	element := listPractice.PushBack("first")

	/*使用list包中提供的方法在句柄后添加元素*/
	listPractice.InsertAfter("second", element) //形参是添加的节点,节点的前一个句柄

	listPractice.InsertBefore("noon", element)

	fmt.Println(listPractice)

	//使用remove删除节点
	listPractice.Remove(element)
	/*
	使用 element 变量，在 element 的位置后面插入 high 字符串。
	使用 element 变量，在 element 的位置前面插入 noon 字符串。
	移除 element 变量对应的元素。
	 */
}