package main

import "fmt"

func main() {
	a := [3]int{1,2,3}
	b := [3]int{4,5,6}
	fmt.Println(a == b, a != b)
}