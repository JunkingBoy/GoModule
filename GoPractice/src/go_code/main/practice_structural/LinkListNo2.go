package main

import (
    "fmt"
)

/*
定义节点
 */
type Node2 struct {
    // 数据域
    Data int
    // 指针域
    next *Node2
}

/*
定义遍历打印方法：
传入指针变量，每次修改指针变量进行读取
 */
func ShowNode2(p *Node2) {
    for p != nil {
        fmt.Println(*p) // 打印指针在此处的地址值
        // 移动指针指向下一个节点
        p = p.next
    }
}

/*
头插法插入链表
 */
func main() {
    // 定义头节点
    var head = new(Node2)
    // 定义头节点数据域
    head.Data = 0

    // 定义头指针，该指针只会指向头节点,这是一个指针变量，指向节点
    var top *Node2
    // 头指针指向头节点，并且始终指向头节点
    top = head
    //
    //// 循环添加节点，当每次添加节点成功了以后修改头指针的指向，让其指向头节点
    //for i := 1; i < 10; i++ {
    //    // 新建节点
    //    var node = Node2{Data: i}
    //    // 每次添加成功都将头指针指向新增的节点
    //    node.next = top
    //    // 给头指针重新赋值
    //    top = &node
    //}

    // 尾插法插入节点数据
    for i := 1; i < 10; i++ {
        // 新建节点
        var node = Node2{Data: i}
        // 每次添加成功都修改头节点的指针域的值，将新插入的节点指针值赋值给头节点
        (*top).next = &node
        // 重新给头节点赋值
        top = &node
    }

    // 调用打印节点方法
    ShowNode2(top)
}
