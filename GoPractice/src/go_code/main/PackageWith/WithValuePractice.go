package main

import (
    "context"
    "fmt"
)

func main() {
    // 定义一个key类型
    type favContextKey string

    // 定义一个变量，从上下文中获取key和value的函数
    f := func(ctx context.Context, k favContextKey) {
        // 判断key值
        if v := ctx.Value(k); v != nil {
            fmt.Println("查询到的值是:", v)
            return
        }
        fmt.Println("找不到键:", k)
    }
    // 创建默认k
    k := favContextKey("language")
    // 创建一个携带key为k，value为"Go"的上下文
    ctx := context.WithValue(context.Background(), k, "Go")

    // 调用f函数
    f(ctx, k)
    f(ctx, favContextKey("color"))
}
