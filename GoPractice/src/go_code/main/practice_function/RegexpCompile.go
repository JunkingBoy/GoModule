package main

import (
    "fmt"
    "regexp"
    "strconv"
)

/*
获取regexp对象，实现匹配、查找、替换等功能
 */
func main() {
    // 提供目标串
    searchIn := "John: 2578.34 William: 4567.23 Steve: 5632.18"
    // 使用正则表达式匹配数字
    pat := "[0-9]+.[0-9]+"

    // 提供函数去匹配字符
    f := func(s string) string {
        // 使用strconv包下提供的函数进行字符匹配
        v, _ := strconv.ParseFloat(s, 32)
        // 返回取出的值
        result := strconv.FormatFloat(v * 2, 'f', 2, 32)
        return result
    }

    // 使用math包下的函数匹配提供的目标串
    if ok, _ := regexp.Match(pat, []byte(searchIn)); ok {
        fmt.Println("找到结果!")
    }

    // 获取compile对象
    re, _ := regexp.Compile(pat)

    // 将匹配部分替换为"**.*"
    str := re.ReplaceAllString(searchIn, "**.*")
    fmt.Println(str)

    // 参数为函数时替换为函数处理的结果
    res := re.ReplaceAllStringFunc(searchIn, f)
    fmt.Println(res)
}
