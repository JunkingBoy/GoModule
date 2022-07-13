package main

/*
创建链表的节点结构体：
1、数据域
2、指针域--->一个指针变量，属于该结构体的指针变量
 */
type Link struct {
	/*创建数据域*/
	element byte
	/*创建指针域--->一个指针变量，属于该结构体*/
	*Link
}

func main() {

}