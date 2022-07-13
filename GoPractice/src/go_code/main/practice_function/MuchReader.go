package main

import (
    "sync"
    "time"
)

/*
定义一个读写锁变量
定义一个读取函数
构建多个协程调用该函数
 */
var m *sync.RWMutex

func moreRead(i int) {
    println(i, "开始读取!")
    // 上锁
    m.RLock()

    // 读取
    println(i, "reading")

    // 休眠
    time.Sleep(1*time.Second)

    // 施放读锁
    m.RUnlock()
    println(i, "读取结束!")
}

// 调用多个读取的函数
func main() {
    m = new(sync.RWMutex)

    // 多个开始读取--->读锁不互斥
    go moreRead(1)
    go moreRead(2)

    // 休眠
    time.Sleep(2*time.Second)
}
