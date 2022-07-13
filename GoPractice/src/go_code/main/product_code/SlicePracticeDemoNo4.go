package main

import "fmt"

func main() {
	a := make([]int, 3)
	b := make([]int, 4, 10)

	fmt.Println(a, b)
	fmt.Println(len(a), len(b))
}