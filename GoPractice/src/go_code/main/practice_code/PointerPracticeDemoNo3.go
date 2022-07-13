package main

import "fmt"

func changePointer(arr, brr *int) (*int, *int) {
	fmt.Printf("Befor change value is:%d,%d\n", arr, brr)
	temp := arr
	arr = brr
	brr = temp
	return brr,arr
}

func main() {
	x, y := 1, 2
	changePointer(&x, &y)
	fmt.Printf("After change value is:%d,%d\n", &x, &y)
	fmt.Println(x, y)
}