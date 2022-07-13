package main

import (
    "fmt"
    "time"
)

func main() {
    // 获取当前时间时间戳
    now := time.Now()

    // 时间戳
    timeStamp1 := now.Unix()

    // 纳秒时间戳
    timestamp2 := now.UnixNano()

    // 现在的时间戳
    fmt.Printf("现在的时间戳: %v\n", timeStamp1)
    // 现在的纳秒时间戳
    fmt.Printf("现在的纳秒时间戳: %v\n", timestamp2)
}
