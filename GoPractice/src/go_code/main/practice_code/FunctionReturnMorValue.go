package main

import "fmt"

func returnMoreValue(arr, brr string) (string, bool) {
	var valueNo1 = arr
	var valueNo2 = brr
	return valueNo1,(valueNo1==valueNo2)
}

func main() {
	var arr string
	var brr string

	fmt.Println("Input string first value :")
	fmt.Scanln(&arr)
	fmt.Println("Input string second value :")
	fmt.Scanln(&brr)

	fmt.Println(returnMoreValue(arr,brr))
}