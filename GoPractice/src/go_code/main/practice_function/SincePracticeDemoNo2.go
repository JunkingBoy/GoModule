package main

import (
	"fmt"
	"time"
)

/*
使用time.now.sub函数进行计算
过程一样不需要变化
 */
func test2() {
	//函数开始时间
	start := time.Now()

	//函数体
	sum := 0
	for i := 0; i < 10000000; i++ {
		sum++
	}

	//函数结束时间
	end := time.Now().Sub(start)
	fmt.Println("函数执行完毕耗时：", end)
}

func main() {
	test2()
}