package main

import "fmt"

func main() {
	a := []int{1,2,3}
	fmt.Println(a[:])
	fmt.Println(a[0:0])
	fmt.Println(a)
	fmt.Println(a[1:2])
}