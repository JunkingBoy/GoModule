package main

import (
    "sync"
    "time"
)

/*
声明一个读写锁变量
定义一个读函数
定义一个写函数
每个函数都是先开始进行读取(写入)
上锁
读取(写入)内容
释放锁
打印结果
 */
var variableM *sync.RWMutex

// 实际调用
func main() {
    variableM = new(sync.RWMutex)
    // 写入的时候与读取或者写入互斥
    go writing(1)
    go reading(2)
    go writing(3)
    // 等待时间
    time.Sleep(2*time.Second)
}

// 读函数
func reading(i int) {
    println(i, "开始读取!")
    // 上锁
    variableM.RLock()
    // 读取内容
    println(i, "读取中...")
    // 等待
    time.Sleep(1*time.Second)
    // 释放锁
    variableM.RUnlock()
    // 打印结果
    println(i, "读取结束!")
}

// 写函数
func writing(i int) {
    println(i, "开始写入!")
    // 上锁
    variableM.Lock()
    // 写入内容
    println(i, "写入中...")
    // 等待
    time.Sleep(1*time.Second)
    // 释放锁
    variableM.Unlock()
    // 打印结果
    println(i, "写入结束")
}
