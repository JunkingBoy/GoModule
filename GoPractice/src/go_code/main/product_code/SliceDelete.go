package main

import "fmt"

func main() {
	/*声明一个切片*/
	var array = []int{1,2,3,4,5,6}
	/*打印出这时候的切片长度*/
	fmt.Println(cap(array))

	array = array[2:] //保留数组2号位后的元素然后复制给原数组，可以利用cap函数看看容量的变化

	fmt.Println(cap(array), array)
}