package main

import "fmt"

func main() {
	var array [2][2]int
	var arrays [2][2]int
	arrays[0][0] = 10
	arrays[0][1] = 20
	arrays[1][0] = 30
	arrays[1][1] = 40
	array = arrays
	for i, j := range array {
		fmt.Printf("第%d位的值是%d\n", i, j)
	}
	var array2 [2]int = arrays[0]
	var value int = arrays[0][1]
	for j, k := range array2 {
		fmt.Printf("第%d位的值是%d\n", j, k)
	}
	fmt.Println(value)
}