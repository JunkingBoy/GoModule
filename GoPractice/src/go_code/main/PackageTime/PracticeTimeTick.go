package main

import (
    "fmt"
    "time"
)

/*
调用时间间隔函数，进行定时器的设置
 */
func main() {
    // 声明一个定时器(一秒一次)
    ticker := time.Tick(time.Second)
    // 循环打印定时器当中的内容
    for i := range ticker {
        fmt.Println(i)
    }
}
