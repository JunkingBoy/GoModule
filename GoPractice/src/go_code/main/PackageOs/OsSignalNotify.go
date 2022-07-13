package main

import (
    "fmt"
    "os"
    "os/signal"
)

/*
使用os/signal包下的函数进行监听
 */
func main() {
    c := make(chan os.Signal, 0)
    // 监听新建的变量
    signal.Notify(c)

    // Block until a signal is received--->阻塞直到收到信号
    s := <-c
    // 打印信号
    fmt.Println("获取到的信号是:", s)
}
