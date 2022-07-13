package main

import "fmt"

func main() {
	var array [4][3]int
	var arrays [2][2]int
	array = [4][3]int{{10, 11}, {20, 21}, {30, 31}, {40, 41}}
	fmt.Println(array)
	array = [4][3]int{1:{0, 21}, 3:{1:41}}
	fmt.Println(array)
	arrays[0][0] = 10
	arrays[0][1] = 20
	arrays[1][0] = 30
	arrays[1][1] = 40
	for i, j := range arrays {
		fmt.Printf("第%d位的值是%d\n", i, j)
	}
}