package main

import (
    "fmt"
    "time"
)

func main() {
    // 获取当前时间
    now := time.Now()

    // 调用time包下的sub函数进行时间截取
    after := now.Sub(now.Local())

    // 打印结果
    fmt.Println(after)
}
