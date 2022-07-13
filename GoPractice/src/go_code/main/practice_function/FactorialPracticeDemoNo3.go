package main

import (
	"errors"
	"fmt"
)

//自定义一个错误
var errorZero = errors.New("division by zero")

/*
定义一个除法函数，被除数除以除数。如果商为0返回错误
 */
func div(dividend, divisor int) (int, error) {
	//判断除数为0抛出异常
	if divisor == 0 {
		return 0, errorZero
	}

	//正常计算，返回空值
	return dividend/divisor, nil //正常计算返回空的error
}

func main() {
	fmt.Println(div(1,0))
}