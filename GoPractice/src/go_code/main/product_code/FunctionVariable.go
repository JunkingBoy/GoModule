package main

import "fmt"

func printValue() {
	fmt.Println("Hello")
}

func main() {
	var f func()

	f = printValue

	f()
}