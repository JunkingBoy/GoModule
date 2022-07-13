package main

import (
    "fmt"
    "sync"
    "time"
)

/*
开启两个协程
创建一个结构体对象
使用两个协程对该变量进行修改
 */
func main() {
    // 声明一个结构体对象，里面存放一个值
    ch := make(chan struct{}, 2)

    // 声明一个互斥锁变量
    var l sync.Mutex

    // 协程一为锁定协程
    go func() {
        // 开始锁
        l.Lock()
        // 释放锁
        defer l.Unlock()
        fmt.Println("goroutine1: 我会锁定2s")
        // 休眠两秒
        time.Sleep(time.Second)
        fmt.Println("goroutine1: 一已解锁!")
        ch <- struct{}{}
    }()

    // 协程二为等待争抢协程
    go func() {
        fmt.Println("goroutine: 等待解锁")
        // 上锁
        l.Lock()
        // 最后解锁
        defer l.Unlock()
        fmt.Println("goroutine: 二已解锁!")
        // 最后关闭将变量放回协程
        ch <- struct{}{}
    }()

    // 等待goroutine执行结束
    for i := 0; i < 2; i++ {
        <-ch
    }
}
