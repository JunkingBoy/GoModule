package main

import "fmt"

func main() {
	switch {
	case true:
		fmt.Println("1、case条件语句为false!")
		fallthrough
	case false:
		fmt.Println("2、case条件语句为true!")
	default:
		fmt.Println("默认的case")
	}
}