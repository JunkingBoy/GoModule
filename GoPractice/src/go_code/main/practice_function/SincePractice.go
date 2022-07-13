package main

import (
	"fmt"
	"time"
)

/*
计算运行时间的函数。使用Since函数，记录函数运行开始时间--->time.Now()
使用Since函数进行时间计算--->time.Since(start)
 */
func test() {
	//函数开始计时器
	start := time.Now()

	//函数体
	sum := 0
	for i := 0; i < 1000000; i++ {
		sum++
	}

	//函数完成时间
	end := time.Since(start)
	fmt.Println("函数执行完毕耗时：", end)
}

func main() {
	test()
}