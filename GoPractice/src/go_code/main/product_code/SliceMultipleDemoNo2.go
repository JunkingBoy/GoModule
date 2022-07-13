package main

import "fmt"

func main() {
	/*声明一个二维切片*/
	slice := [][]int{{10}, {100, 200}}

	fmt.Println(len(slice))
	fmt.Println(cap(slice))

	/*再第二维度的零号索引处添加一个元素30*/
	slice[0] = append(slice[0], 30) //这里的索引位置要当作数组的索引位置来看待

	fmt.Println(slice)
	fmt.Println(len(slice))
	fmt.Println(cap(slice))
}