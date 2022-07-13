package main

import (
    "flag"
    "fmt"
)

/*
调用parse函数对命令行参数进行解析
 */
// 定义命令行参数
var Input_pstrName = flag.String("name", "Lucifer", "input user name")
var Input_piAge = flag.Int("age", 22, "input user age")
var Input_flagvar int

// 定义一个初始化函数，返回一个参数变量指针
func Init() {
    flag.IntVar(&Input_flagvar, "flagname", 1234, "参数帮助信息!")
}

// 调用parse函数解析命令行参数
func main() {
    Init()
    flag.Parse()
    fmt.Printf("args=%s, num=%d\n", flag.Args(), flag.NArg())
    for i := 0; i != flag.NArg(); i++ {
        fmt.Printf("arg[%d]=%s\n", i, flag.Arg(i))
    }
    fmt.Println("name=", *Input_pstrName)
    fmt.Println("age=", *Input_piAge)
    fmt.Println("flagname=", Input_flagvar)
}
