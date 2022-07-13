package main

import (
    "fmt"
    "time"
)

func main() {
    // 获取当前时间
    now := time.Now()
    // 序列化输出结果
    fmt.Printf("当前时间为:%v\n", now)

    // 通过包下提供的函数获取年\月\日\时\分\秒等信息
    year := now.Year()
    month := now.Month()
    day := now.Day()
    hour := now.Hour()
    minute := now.Minute()
    second := now.Second()

    // 序列化输出结果集
    fmt.Printf("%d-%02d-%02d %02d:%02d:%02d\n", year, month, day, hour, minute, second)
}
