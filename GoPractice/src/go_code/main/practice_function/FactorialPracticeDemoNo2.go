package main

import (
	"errors"
	"fmt"
)

var (
	err = errors.New("This is an error!!!")
)

func main() {
	fmt.Printf("%d is A: is %t\n", 16, A(16))
	fmt.Printf("%d is B: is %t\n", 17, B(17))
	fmt.Printf("%d is C: is %t\n", 18, B(18))
}

func A(n int) bool {
	if n == 0 {
		return true
	}
	return B(C(n)-1)
}

func B(n int) bool {
	if n == 0 {
		return false
	}
	return A(C(n)-1)
}

func C(n int) int {
	if n < 0 {
		return -n
	}
	return n
}