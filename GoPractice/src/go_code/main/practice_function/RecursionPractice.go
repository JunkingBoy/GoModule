package main

import "fmt"

func main() {
	result := 0
	//循环传递n
	for i := 0; i <= 10; i++ {
		result = fibonacci(i)
		fmt.Printf("fibonacci(%d) is %d\n", i, result)
	}
}

func fibonacci(n int) (res int) {
	if n < 2 {
		res = 1
	}else {
		res = fibonacci(n-1) + fibonacci(n-2)
	}
	return
	/*
	这个return只是结束条件
	 */
}