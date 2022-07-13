package main

import "fmt"

func big(arr, brr int) bool {
	if (arr - brr) > 0 {
		return true
	}else {
		return false
	}
}

func main() {
	if err := big(3, 5); err != false {
		fmt.Println(true)
		return
	}else {
		fmt.Println(false)
		return
	}
}