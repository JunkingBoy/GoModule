package main

import "fmt"

func main() {
	/*声明一个数组*/
	var num [10]int

	/*循环添加值*/
	for i := 0; i < len(num); i++ {
		num[i] = (i+1)
	}

	/*循环打印出数组的value*/
	for _, value := range num {
		fmt.Println("Value is :", value)
	}
}