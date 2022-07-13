package main

import (
    "fmt"
    "time"
)

func main() {
    // 获取当前时间
    t := time.Now()

    u := t.Unix()

    // 打印当前时间
    fmt.Println(u)

    // 使用time包提供的函数判断当前时间戳是周几
    fmt.Println(t.Weekday().String())
}
