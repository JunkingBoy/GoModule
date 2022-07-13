package main

import (
    "context"
    "fmt"
    "time"
)

func main() {
    // 设置一个超时的deadline
    d := time.Now().Add(50 * time.Millisecond)
    // 调用withdeadline函数在超过时间以后结束goroutine
    ctx, cancel := context.WithDeadline(context.Background(), d)

    // 尽管ctx会过期，但在任何情况下调用它的cancel函数都是很好的实践。
    // 如果不这样做，可能会使上下文及其父类存活的时间超过必要的时间。
    /* 执行cancel函数 */
    defer cancel()

    // 使用select选择器根据情况执行代码
    select {
    case <-time.After(1 * time.Second):
        fmt.Println("overslept")
    case <-ctx.Done():
        fmt.Println(ctx.Err())
    }
}
