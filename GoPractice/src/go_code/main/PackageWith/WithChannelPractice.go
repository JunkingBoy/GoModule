package main

import (
    "context"
    "fmt"
)

/*
设计一个`gen`函数，在单独的`Goroutine`中生成整数并将它们发送到返回的通道。
`gen`的调用者在使用生成的整数之后要取消上下文，以免`gen`启动的内部`Goroutine`发生泄漏
 */
func main() {
    gen := func(ctx context.Context) <-chan int {
        // 定义整数
        dst := make(chan int)
        n := 1
        // 开启一个goroutine
        go func() {
            // 循环生成整数并发送到返回的通道
            for {
                select {
                case <-ctx.Done():
                    // 结束该routine,防止泄露
                    return
                case dst <- n:
                    n++
                }
            }
        }()
        return dst
    }

    // 调用withchannel函数关闭通道
    ctx, cancel := context.WithCancel(context.Background())
    // 取完需要的整数后调用cancel函数
    defer cancel()

    for n := range gen(ctx) {
        fmt.Println(n)
        if n == 5 {
            break
        }
    }
}
