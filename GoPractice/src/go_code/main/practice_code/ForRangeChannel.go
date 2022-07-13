package main

import "fmt"

func main() {
	/*声明一个通道,放如整型数据*/
	c := make(chan int)

	go func() {

		c <- 1
		c <- 2
		c <- 3
		close(c)
	}()

	/*循环输出通道的值*/
	for value := range c {
		fmt.Println(value)
	}
}