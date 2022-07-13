package main

import (
    "errors"
    "fmt"
    "math"
)

/* 创建error接口 */
func Sqrt(f float64) (float64, error) {
    if f < 0 {
        return -1, errors.New("方法报错!")
    }
    return math.Sqrt(f), nil
}

/* 方法调用 */
func main() {
    result, err := Sqrt(-13)
    if err != nil {
        fmt.Println(err)
    }else {
        fmt.Println(result)
    }
}
