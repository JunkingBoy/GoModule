package main

import "fmt"

var (
	cirCula bool
)

func main() {
	cirCula = true
	for cirCula == true {
		fmt.Println("这是一个真值，其值为:", cirCula)
		break
	}
	fmt.Println("这不是一个真值，其值为:", cirCula)
}