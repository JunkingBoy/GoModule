package main

import (
	"fmt"
	"time"
)

func main() {
	//构建一个通道
	chNo2 := make(chan int)

	//开启一个并发匿名函数
	go func() {

		//for循环，从3->0
		for i := 3; i > 0; i-- {
			//每次将循环的值发送到main的goroutine中
			chNo2 <- i

			//每次发送完时等待
			time.Sleep(time.Second)
		}
	}()

	//遍历接收通道的数据
	for data := range chNo2 {
		//打印通道数据
		fmt.Println(data)

		//当遇到数据0时，退出接收循环
		if data == 0 {
			break
		}
	}
}