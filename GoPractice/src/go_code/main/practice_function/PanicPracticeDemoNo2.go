package main

import "fmt"

func main() {
	defer fmt.Println("宕机后发生的事情1")
	defer fmt.Println("宕机后发生的事情2")

	//手动触发宕机
	panic("宕机")

	fmt.Println("事情3")
}