package main

import "fmt"

//递归实现
func Fibo1(n int) int {
	if n == 0 {
		return 0
	}else if n == 1 {
		return 1
	}else if n > 1 {
		return Fibo1(n-1) + Fibo1(n-2)
	}else {
		return -1
	}
}

func Fibo2(n int) int {
	if n < 0 {
		return -1
	}else if n == 0 {
		return 0
	}else if n <= 2 {
		return 1
	}else {
		fibolNo1, fibolNo2 := 1, 1
		result := 0
		for i := 3; i <= n; i++ {
			result = fibolNo1 + fibolNo2
			fibolNo1, fibolNo2 = fibolNo2, result
		}
		return result
	}
}

func main() {
	n := 7
	i := Fibo1(n)
	j := Fibo2(n)
	fmt.Println(i)
	fmt.Println(j)
}