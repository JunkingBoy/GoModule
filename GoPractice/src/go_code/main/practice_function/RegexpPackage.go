package main

import (
    "fmt"
    "regexp"
)

/*
使用regexp包下的函数进行指定类型字符串的匹配
 */
func main() {
    buf := "abc azc a7c aac 888 a9c  tac"

    /* 解析正则表达式，如果成功，返回解析器 */
    regeNo1 := regexp.MustCompile(`a[0-9]c`)
    // 为空，返回异常
    if regeNo1 == nil {
        fmt.Println("regexp error")
        return
    }

    // 根据正则规则提取关键信息
    result1 := regeNo1.FindAllStringSubmatch(buf, -1)
    fmt.Println("结果为：",  result1)
}
