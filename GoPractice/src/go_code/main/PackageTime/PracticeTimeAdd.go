package main

import (
    "fmt"
    "time"
)

/*
调用time包下的函数求时间的加法
 */
func main() {
    now := time.Now()
    // 调用time包下的add函数求出一小时后的时间
    later := now.Add(time.Hour)
    // 打印结果
    fmt.Println(later)
}
