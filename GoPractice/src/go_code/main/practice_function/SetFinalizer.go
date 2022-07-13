package main

import (
    "log"
    "runtime"
    "time"
)

/*
定义终止器去自动的执行垃圾回收
 */
// 定义一个类型(相当于类)
type Road int

// 定义一个方法，需要传入类型的指针
func findRoad(r *Road) {
    // 调用log包下的方法打印指针
    log.Println("Road:", *r)
}

// 定义一个函数，定义为终止器
func entry() {
    // 定义一个Road对象
    var rd Road = Road(999)
    // 获取变量指针
    r := &rd

    // 设置终止器
    runtime.SetFinalizer(r, findRoad)
}

// 调用定义的方法
func main() {
    entry()

    for i := 0; i < 10; i++ {
        time.Sleep(time.Second)
        runtime.GC()
    }
}
