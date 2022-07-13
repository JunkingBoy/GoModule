package main

import "fmt"

func main() {
	str := new(string) //有指向的地址值

	var str2 *string //指向nil

	fmt.Println(str)
	fmt.Println(str2)
}