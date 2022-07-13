package main

import "fmt"

/* 打印类型的方法--->使用类型断言 */
func printType(v interface{}) {
    switch v.(type) {
    case int:
        fmt.Println(v, "is int!")
    case string:
        fmt.Println(v, "is string!")
    case bool:
        fmt.Println(v, "is boolean!")
    default:
        fmt.Println(v, "is unknowtype!")
    }
}

/* 调用函数 */
func main() {
    printType(1024)
    printType("HelloWorld!")
    printType(false)
}
