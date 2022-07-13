package main

import "fmt"

func main() {
	var a  = [3]int{1, 2, 3}
	fmt.Println(a, a[1:2])
	fmt.Println(a[1:])
	fmt.Println(a[:2])
	fmt.Println(a[:])
}