package main

import "fmt"

/*
定义节点
 */
type Node struct {
    // 数据域
    Data int

    // 指针域
    Next *Node

    /*
    Data 用来存放结点中的有用数据
    Next 是指针类型的成员，它指向 Node struct 类型数据，也就是下一个结点的数据类型
     */
}

/*
定义一个遍历链表节点的方法
 */
func ShowNode(p *Node) {
    // 节点不为空就打印
    for p != nil {
        if p != nil {
            fmt.Println(*p) // '*'符号+变量名获取到变量的指针地址
            // 指针域指向下一个节点的数据域
            p = p.Next
        }
    }
}

/*
调用该方法
 */
func main() {
    // 定义头节点
    var head = new(Node)
    // 定义该节点的数据域和指针域
    head.Data = 1

    // 定义第二个节点
    var node1 = new(Node)
    // 定义数据域
    node1.Data = 2

    // 讲头节点指向第二个节点
    head.Next = node1

    // 定义第三个节点
    var node2 = new(Node)
    // 定义数据域
    node2.Data = 3

    // 将第二个节点与第三个节点连接
    node1.Next = node2

    //调用打印链表的方法
    ShowNode(head)
}
