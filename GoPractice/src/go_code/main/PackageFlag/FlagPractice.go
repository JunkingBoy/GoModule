package main

import (
    "flag"
    "fmt"
    "strings"
)

// 定义一个类型，用于增加该类型的方法
type sliceValue []string

// 定义一个创建该类型的函数--->返回该类型的指针
func newSliceValue(vals []string, p *[]string) *sliceValue {
    *p = vals
    // 返回指针--->转型
    return (*sliceValue)(p)
}

// 定义一个value接口，里面实现的获取字符串和设置字符串的函数
type Value interface {
    String() string
    Set(string) error
}

// sliceValue类型实现flag包中的Value接口，将命令行接收到的值用,分隔存到slice里
func (s *sliceValue) Set(val string) error {
    // 用逗号分割存到slice中--->*+变量是获取指针类型变量的具体变量
    *s = sliceValue(strings.Split(val, ","))
    return nil
}

// 实现set函数
func (s *sliceValue) String() string {
    *s = sliceValue(strings.Split("我全都要", ","))
    return "It's none of my business"
}

/*
可执行文件名 -slice="java,go"  最后将输出[java,go]
可执行文件名 最后将输出[default is me]
*/
// 调用函数
func main() {
    // 定义一个切片变量
    var languages []string
    // 使用flag.Var函数自定义flag
    flag.Var(newSliceValue([]string{}, &languages), "slice", "我喜欢的语言是:")
    // 解析参数
    flag.Parse()

    // 打印切片变量
    fmt.Println(languages)
}
