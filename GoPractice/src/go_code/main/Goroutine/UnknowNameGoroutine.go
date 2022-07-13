package main

import (
    "fmt"
    "time"
)

/*
使用匿名参数的方法创建一个goroutine
 */
func main() {
    // 声明一个匿名goroutine
    /*
    没有调用参数，所以再启动main的时候会直接被执行
     */
    go func() {
        var times int

        // 无线添加
        for {
           times++
           fmt.Println("tick", times)
           // 等待
           time.Sleep(time.Second)
        }
    }()

    // 接收命令行参数
    var input string
    fmt.Scanln(&input)
}
