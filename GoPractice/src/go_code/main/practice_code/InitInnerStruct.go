package main

import "fmt"

type WheelNo2 struct {
    size int
}

/*
该结构体内内嵌一个匿名结构体
 */
type CarNo2 struct {
    WheelNo2
    //引擎
    EngineNo2 struct{
        Power int
        Type string
    }
}

func main() {
    //初始化最外层结构体变量
    c := CarNo2{
        WheelNo2:WheelNo2{
            22,
        },
        //因为EngineNo2是匿名内嵌结构体，所以使用初始化匿名结构体的方法
        EngineNo2: struct {
            Power int
            Type  string
        }{Power: 100, Type: "大功率"},
    }

    fmt.Println(c)
}