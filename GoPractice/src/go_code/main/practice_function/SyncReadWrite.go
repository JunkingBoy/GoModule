package main

import (
    "fmt"
    "math/rand"
    "sync"
)

/*
开启两个协程
一个读
一个写
同时对结构体当中的属性进行操作
体会读写锁的特性
 */
// 声明两个变量，一个int的总数变量，一个读写锁变量
var count int
var rw sync.RWMutex

func main() {
    // 设置结构体变量
    ch := make(chan struct{}, 10)

    // 循环进行读或者写的操作
    for i := 0; i < 5; i++ {
        // 调用read方法
        go read(i, ch)
    }

    // 循环进行写的操作
    for i := 0; i < 5; i++ {
        // 调用write方法
        go write(i, ch)
    }

    // 一次将值放回
    for i := 0; i < 10; i++ {
        <-ch
    }
}

/* 构造读取函数 */
func read(n int, ch chan struct{}) {
    // 设置读锁定
    rw.RLock()
    fmt.Printf("goroutine %d 进入读操作...\n", n)
    v := count
    fmt.Printf("goroutine %d 读取结束，值为：%d\n", n, v)
    // 解锁
    rw.RUnlock()
    // 放回
    ch <- struct{}{}
}

/* 构造写入函数 */
func write(n int, ch chan struct{}) {
    // 设置写锁
    rw.Lock()
    // 开始写入
    fmt.Printf("goroutine %d 进入写操作...\n", n)
    // 设置随机数
    v := rand.Intn(1000)
    fmt.Printf("goroutine %d 写入结束， 新的值为：%d\n", n, v)
    // 释放写锁
    rw.Unlock()
    // 放回
    ch <- struct{}{}
}
