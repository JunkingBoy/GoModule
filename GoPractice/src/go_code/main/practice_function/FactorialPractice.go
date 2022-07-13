package main

import "fmt"

/*
阶乘的函数，传入一个无符号的整数形参，返回的也是一个无符号的整数
 */
func factorial(n uint64) (result uint64) {
	if n > 0 {
		result = n*factorial(n-1)
		return result
	}
	return 1
}

func main() {
	fmt.Println(factorial(3))
}