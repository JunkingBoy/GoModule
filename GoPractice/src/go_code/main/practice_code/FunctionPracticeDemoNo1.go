package main

import "fmt"

func max(num1, num2 int) int {
	var result int
	if num1-num2 > 0 {
		result = num1
		return result
	}else {
		result = num2
		return result
	}
}

func main() {
	var valueNum1 int
	var valueNum2 int
	var resultValu1 int

	//键盘输入
	fmt.Println("Input first value:")
	fmt.Scanln(&valueNum1)
	fmt.Println("Input Second value:")
	fmt.Scanln(&valueNum2)

	resultValu1 = max(valueNum1, valueNum2)

	fmt.Println("The max value is :", resultValu1)
}