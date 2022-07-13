package main

import "fmt"

/*
定义双向链表节点信息
 */
type Node4 struct {
    // 前驱节点
    prev *Node4

    // 数据域
    Data int

    // 后继节点
    next *Node4
}

/*
遍历节点的方法
 */
func ShowNode4(p *Node4) {
    for p != nil {
        fmt.Println(*p)
        // 打印前驱节点如果存在
        if p.prev != nil {
            fmt.Println(*p.prev)
        }
        // 将指针移向下一个节点
        p = p.next
    }
}

/*
遍历双向链表
 */
func main() {
    // 定义头节点
    var head = new(Node4)
    // 定义头节点的数据域
    head.Data = 0
    // 定义头节点的指针域
    head.prev = nil

    // 定义头指针
    var headTop *Node4

    // 添加节点使用尾插法
    for i := 1; i < 10; i++ {
        // 新建节点
        var node = Node4{Data: i}
        // 修改头指针指向新建节点
        (*headTop).next = &node
        // 修改新建的节点的前驱指针指向头指针
        node.prev = headTop
        // 更新头指针
        headTop = &node
    }

    ShowNode4(headTop)
}
