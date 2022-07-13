package main

import (
	"fmt"
	"time"
)

/*
实现一个斐波那契数列，使用递归的方法
 */
func fib(n int) (res int) {
	if n <= 2 {
		res = 1
	}else {
		res = fib(n-1) + fib(n-2)
	}
	return //结束方法
}

func main() {
	result := 0
	start := time.Now()
	for i := 1; i <= 40; i++ {
		result = fib(i)
		fmt.Printf("数列第 %d 位: %d\n", i, result)
	}
	end := time.Now()
	delta := end.Sub(start)
	fmt.Printf("程序的执行时间为: %s\n", delta)
}