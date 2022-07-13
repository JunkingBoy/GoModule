package main

import "fmt"

func main() {
	var array [30]int
	for i := 0; i < 30; i++ {
		array[i] = i+1
	}

	//区间内的值
	fmt.Println(array[10:20])

	//中间到尾部的所有元素
	fmt.Println(array[15:])

	//开头到中间指定位置的值
	fmt.Println(array[:15])
}