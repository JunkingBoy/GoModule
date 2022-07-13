package main

import (
    "fmt"
    "time"
)

func main() {
    // 获取当前时区时间
    t := time.Now()

    // 调用equal函数比较时间
    bool := t.Equal(time.Now())
    fmt.Println(bool)
}
