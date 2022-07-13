package main

import (
    "fmt"
    "math"
)

/* 定义错误类型 */
type definedError struct {
    Num float64
    problem string
}

/* 定义类型实现的接口 */
func (e definedError) Error() string {
    return fmt.Sprintf("错误，原因是： \"%f\"是个自然数", e.Num)
}

/* 创建error接口 */
func Sqrt1(f float64) (float64, error) {
    if f < 0 {
        // 使用自定义错误类型进行返回
        return -1, definedError{Num: f}
    }
    return math.Sqrt(f), nil
}

/* 方法调用 */
func main() {
    result, err := Sqrt1(-13)
    if err != nil {
        fmt.Println(err)
    }else {
        fmt.Println(result)
    }
}
