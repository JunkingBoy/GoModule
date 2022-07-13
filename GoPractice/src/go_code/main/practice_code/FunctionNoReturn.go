package main

import "fmt"

func noReturn(arr int) {
	fmt.Println("Value is :", arr)
}

func main() {
	var input int
	fmt.Println("Input value :")
	fmt.Scanln(&input)
	noReturn(input)
}