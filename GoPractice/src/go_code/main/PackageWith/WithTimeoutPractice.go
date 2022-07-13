package main

import (
    "context"
    "fmt"
    "time"
)

func main() {
    // 传递超时的上下文信息
    ctx, cancel := context.WithTimeout(context.Background(), 50 * time.Microsecond)

    // 告诉阻塞函数在超时结束后应该放弃其工作--->调用取消函数
    defer cancel()

    // 通过select选择执行的函数
    select {
    case <-time.After(1 * time.Second):
        fmt.Println("overslept")
    case <-ctx.Done():
        // 终端输出"context deadline exceeded"
        fmt.Println(ctx.Err())
    }
}
