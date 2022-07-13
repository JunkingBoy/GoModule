package main

import "fmt"

/*
定义三个结构体
 */
type Wheel struct {
    Size int
}

type Engine struct {
    Power int //功率
    Type string //类型
}

//结构体：车
type Car struct {
    //内嵌结构体
    Wheel
    Engine
}

func main() {
    //初始化结构体Car同时初始化内嵌的结构体
    c := Car{
        //吃实话内嵌的结构体，视内嵌的结构体为成员进行初始化
        Wheel:Wheel{
            Size: 18,
        },
        //初始化另一个内嵌结构体，视为成员
        Engine:Engine{
            Power: 100,
            //可以选择成员进行初始化
        },
    }

    //打印c
    fmt.Println(c)
}