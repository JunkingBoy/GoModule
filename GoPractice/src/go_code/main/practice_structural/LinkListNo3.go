package main

import "fmt"

/*
定义节点
 */
type Node3 struct {
    // 数据域
    Data int
    
    //指针域
    next *Node3
}

/*
循环打印指针节点内容方法
 */
func ShowNode3(p *Node3) {
    // 如果p不为空就打印内容
    for p != nil {
        fmt.Println(*p)
        // 移动指针指向下一个节点
        p = p.next
    }
}

/*
双指针方法定位循环链表结构
 */
func main() {
    // 定义头节点
    var head = new(Node3)
    // 赋值头节点的数据域
    head.Data = 0

    // 定义两个指针，一个作为头指针一个作为移动指针
    var headTop *Node3
    var moveNode *Node3

    headTop = head
    moveNode = head

    // 循环添加节点，使用尾插法
    for i := 1; i < 10; i++ {
        // 新建节点
        var node = Node3{Data: i}
        // 每次修改移动指针的指针域指向新建的节点的地址
        (*moveNode).next = &node
        // 重新给移动节点赋值
        moveNode = &node
        if i == 9 {
            // 将移动节点的指针域指向头节点
            (*moveNode).next = headTop
        }
    }

    ShowNode3(headTop)
}
