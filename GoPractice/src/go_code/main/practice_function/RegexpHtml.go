package main

import (
    "fmt"
    "regexp"
)

/*
匹配<div>标签中的内容
 */
func main() {
    // 声明html变量
    buf := `
    <!DOCTYPE html>
<html lang="zh-CN">
<head>
    <meta charset="UTF-8">
    <title>Lucifer's Home</title>
</head>
<body>
    <div>WelCome!</div>
    <div>This is My Home:
        房间A
        房间B
        房间C
    </div>
    <div>洗手间</div>
    <div>卫生间</div>
</body>
</html>
`
    // 使用解析正则表达式--->匹配div当中的内容
    reg := regexp.MustCompile(`<div>(?s:(.*?))</div>`)
    if reg == nil {
        fmt.Println("内容错误!")
        return
    }

    // 提取关键信息
    result := reg.FindAllStringSubmatch(buf, -1)

    // 过滤掉<>和</>
    for _, text := range result {
        fmt.Println("text[1] = ", text[1])
    }
}
