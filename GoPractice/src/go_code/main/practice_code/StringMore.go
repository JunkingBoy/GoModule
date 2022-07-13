package main

import "fmt"

func main() {
	const value = `
		第一行
		第二行
		第三行
		\r\n
	`

	fmt.Println(value)
}