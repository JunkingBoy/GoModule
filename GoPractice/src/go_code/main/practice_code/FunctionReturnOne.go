package main

import "fmt"

func returnValue(arr int) int {
	fmt.Println("Return Value is :", arr)
	return arr
}

func main() {
	var input int
	fmt.Println("Input Value is :")
	fmt.Scanln(&input)
	returnValue(input)
}