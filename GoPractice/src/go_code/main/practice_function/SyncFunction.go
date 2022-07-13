package main

import (
    "fmt"
    "time"
)

/*
设计一个并发场景，不使用锁看看结果如何
 */
func main() {
    // 设置一个变量a，使用两个协程对其值进行累加
    var a = 0
    for i := 0; i < 1000; i++ {
        // 开启一个协程
        go func(index int) {
            a += 1
            fmt.Println(a)
        }(i)
    }

    // 每次累加以后协程暂歇两秒
    time.Sleep(time.Second)
}
