package main

import "fmt"

func main() {
	/*声明一个切片*/
	var array []int

	/*追加一个元素*/
	array = append(array, 1)

	/*追加多个元素*/
	array = append(array, 1,2,3)

	/*追加一个切片*/
	array = append(array, []int{1,2,3,4,5}...)
	/*
	追加一个切片需要对追加的切片进行压缩。书写格式为：[]int{}...
	后面三个"..."符号必不可少，是压缩的标识符
	[]int{1,2,3,4,5}...这个不是一个新建的数组切片对象放入array当中，而是将元素放入slice当中
	 */

	fmt.Println(array)
}