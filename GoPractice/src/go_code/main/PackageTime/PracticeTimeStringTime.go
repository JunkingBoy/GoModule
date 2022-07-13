package main

import (
    "fmt"
    "time"
)

/*
调用time包下的函数解析字符串时间
 */
func main() {
    var layout string = "2006-01-02 15:04:05"
    var timeStr string = "2019-12-12 15:22:12"

    // 调用time包下的函数
    timeObject, _ := time.Parse(layout, timeStr)
    fmt.Println(timeObject)

    // 调用ParseInLocation函数
    timeObject2, _ := time.ParseInLocation(layout, timeStr, time.Local) // time.Local直接获取到当前时区的时间
    fmt.Println(timeObject2)
}
