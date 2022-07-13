package main

import "fmt"

func main() {
	/*声明一个切片*/
	var array = []int{1, 2, 3, 4, 5, 6, 7}

	fmt.Println(array[:2])

	/*使用append方法删除中间部分的元素*/
	//删除掉元素3
	array = append(array[:2], array[3:]...)

	fmt.Println(array)
}