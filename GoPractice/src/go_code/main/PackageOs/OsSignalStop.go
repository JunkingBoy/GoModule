package main

import (
    "fmt"
    "os"
    "os/signal"
)

/*
使用os/signal包下的stop函数取消监听
 */
func main() {
    c := make(chan os.Signal, 0)
    // 调用notify函数开启监听
    signal.Notify(c)

    // 禁止继续往c中存入内容
    signal.Stop(c)
    // c无内容，此处阻塞，所以不会执行下面的语句，也就没有输出
    s := <-c
    fmt.Println("获取到的信号是:", s)
}
