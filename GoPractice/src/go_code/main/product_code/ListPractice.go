package main

import (
	"container/list"
	"fmt"
)

func main() {
	/*初始化一个链表*/
	listPractice := list.New()

	//另一种声明方式
	var listPracticeNo2 list.List

	fmt.Println(listPractice, listPracticeNo2)
}