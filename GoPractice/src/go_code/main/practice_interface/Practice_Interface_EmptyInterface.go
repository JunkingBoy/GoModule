package main

import "fmt"

func main() {
    // 声明一个空接口
    var any interface{}

    // 给空接口赋值
    any = 1
    // 提供给 fmt.Println 的类型依然是 interface{}
    fmt.Println(any)

    // 为 any 赋值一个字符串 hello。此时 any 内部保存了一个字符串。但类型依然是 interface{}
    any = "hello"
    fmt.Println(any)

    any = false
    fmt.Println(any)
}
