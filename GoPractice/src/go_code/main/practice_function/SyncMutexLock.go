package main

import (
    "fmt"
    "sync"
    "time"
)

/*
设置一个累加变量
使用互斥锁对变量进行锁定控制
*/
func main() {
    // 累加变量
    var a = 0

    // 锁变量
    var lock sync.Mutex

    // 循环累加变量
    for i := 0; i < 1000; i++ {
        // 开启一个协程
        go func(index int) {
            // 进行加锁操作
            lock.Lock()

            // 最后的处理是释放掉锁
            defer lock.Unlock()

            // 进行累加
            a += 1
            fmt.Printf("goruntine %d, a = %d\n", index, a)
        }(i)
    }

    // 等待一秒结束主程序
    // 确保所有协程执行完成
    time.Sleep(time.Second)
}
