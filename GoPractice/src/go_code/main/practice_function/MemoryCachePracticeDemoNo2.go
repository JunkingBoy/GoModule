package main

import (
	"fmt"
	"time"
)

/*
先预声明一块内存区域，大小比fib数列传入的形参要大
 */
const LIM = 41

var fibs [LIM]uint64 //提前声明的一块内存区域

/*
实现fib数列，每次运算之前查询一下内存区域的
 */
func fibonacci2(n int) (res uint64) {
	//记忆化：检查数组中是否已知斐波那契（n）
	//查询内存缓存
	if fibs[n] != 0 {
		/*fibs赋值给res，下次计算的时候使用res进行计算*/
		res = fibs[n]
		return
	}
	if n <= 2 {
		res = 1
	}else {
		res = fibonacci2(n-1) + fibonacci2(n-2)
	}
	fibs[n] = res
	return
}

func main() {
	var result uint64 = 0
	start := time.Now()
	for i := 1; i < LIM; i++ {
		result = fibonacci2(i)
		fmt.Printf("数列第 %d 位: %d\n", i, result)
	}
	end := time.Now()
	delta := end.Sub(start)
	fmt.Printf("程序的执行时间为: %s\n", delta)
}