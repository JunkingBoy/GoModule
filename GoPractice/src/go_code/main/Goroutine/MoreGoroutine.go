package main

import (
    "fmt"
    "runtime"
    "sync"
)

/* 声明一个counter变量 */
var counter int = 0

/* 调用sync包下的函数 */
func Count(lock *sync.Mutex) {
    /* 上锁 */
    lock.Lock()
    /* 变量自增，防止争抢 */
    counter++
    // 打印结果
    fmt.Println(counter)
    // 解锁
    lock.Unlock()
}

/* 调用main函数 */
func main() {
    /* 声明sync.Mutex结构体对象 */
    lock := &sync.Mutex{}
    // 循环开启十个goroutine调用Count函数
    for i := 0; i < 10; i++ {
        /*
        特点：
        1、10 个 goroutine 中共享了变量 counter。每个 goroutine 执行完成后，会将 counter 的值加 1
        2、10 个 goroutine 是并发执行的，引入了锁，每次对变量的操作都需要先锁住，操作完成后将锁释放
         */
        go Count(lock)
    }

    /* 使用 for 循环来不断检查 counter 的值（同样需要加锁）。当其值达到 10 时，说明所有 goroutine 都执行完毕了，这时主函数返回，程序退出。 */
    for {
       lock.Lock()
       /* 检查counter变量所以声明count变量 */
       c := counter
       lock.Unlock()
       /* 判断如果是到10了直接退出 */
       runtime.Gosched()
        if c >= 10 {
            break
        }
    }
}
