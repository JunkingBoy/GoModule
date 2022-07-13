package main

import "fmt"

func main() {
	a := 60
	b := 13
	c := 0

	c = a & b
	fmt.Println("第一行 -c 的值为 %d\n", c)

	c = a | b
	fmt.Println("第二行 -c 的值为 %d\n", c)

	c = a ^ b
	fmt.Println("第三行 -c 的值为 %d\n", c)

	c = a << 2
	fmt.Println("第四行 -c 的值为 %d\n", c)

	c = a >> 2
	fmt.Println("第五行 -c 的值为 %d\n", c)
}