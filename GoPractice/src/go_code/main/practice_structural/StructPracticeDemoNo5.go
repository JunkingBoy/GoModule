package main

import "fmt"

/*
创建一个省份结构体
 */
type Address struct {
	Province string
	City string
	Code int64
}

func main() {
	address := Address{
		"广西",
		"柳州",
		545001,
	}
	fmt.Println(address)
}