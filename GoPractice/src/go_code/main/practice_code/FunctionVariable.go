package main

import (
	"fmt"
	"math"
)

func main() {
	/*声明函数变量*/
	getSquareRoot := func(num float64) float64 {
		return math.Sqrt(num)
	}

	/*使用函数*/
	fmt.Println(getSquareRoot(8))
}