package main

import (
    "fmt"
    "time"
)

func main() {
    // 获取当前时间
    now := time.Now()
    // 格式化的模板为Go的出生时间2006年1月2号15点04分 Mon Jan
    // 24h
    fmt.Println(now.Format("2006-01-02 15:04:05.000 Mon Jan"))
    // 12h
    fmt.Println(now.Format("2006-01-02 03:04:05.000 PM Mon Jan"))
    fmt.Println(now.Format("2006/01/02 15:04"))
    fmt.Println(now.Format("15:04 2006/01/02"))
    fmt.Println(now.Format("2006/01/02"))
}
