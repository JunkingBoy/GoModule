package main

import "fmt"

func main() {
    // 声明接口变量
    var x interface{}
    // 赋值
    x = 10
    //// 使用接口断言判断(判断接口当中的值和类型满不满足条件)
    //value, ok := x.(int)
    //fmt.Println(value, ",", ok)

    // 调用类型判断方法进行判断
    getType(x)
}

/*
定义一个接口类型判断的函数。根据传参判断接口的参数类型
 */
func getType(a interface{}) {
    switch a.(type) {
    case int:
        fmt.Println("入参类型为int类型!")
    case string:
        fmt.Println("入参的类型为string类型!")
    case float64:
        fmt.Println("入参的类型为float类型!")
    default:
        fmt.Println("未知类型!")
    }
}
