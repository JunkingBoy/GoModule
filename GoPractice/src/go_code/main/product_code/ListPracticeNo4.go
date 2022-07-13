package main

import (
	"container/list"
	"fmt"
)

func main() {
	/*声明一个列表*/
	listPractice := list.New()

	//尾部添加
	listPractice.PushBack("first")
	//头部添加
	listPractice.PushFront(77)

	//循环遍历--->注意格式
	for i := listPractice.Front(); i != nil; i = i.Next() {
		fmt.Println(i.Value)
	}
}