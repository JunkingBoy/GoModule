package main

import (
    "fmt"
    "time"
)

/*
定义runnning函数，同时定义接收命令行输入函数
并发执行
 */
func runnnig() {
    var times int
    // 无限循环添加
    for {
       times++
       // 打印结果
       fmt.Println("answer:", times)
       // 延时一秒
       time.Sleep(time.Second)
    }
}

// 并发执行
func main() {
    // 定义并发执行程序
    go runnnig()

    // 命令行接收输入
    var input string
    fmt.Scanln(&input)
}
