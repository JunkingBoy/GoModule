package main

import "fmt"

/*
手写一个交换函数，实现传入两个数，交换他们的值
 */
func swap(arr int, brr int) (int, int) { //函数调用前己经为返回值和参数分配了栈空间，分配顺序是从右向左的，先是返回值，然后是参数
	/*交换值*/
	arr, brr = brr, arr
	return arr, brr
}

func main() {
	fmt.Println(swap(10, 20))
}