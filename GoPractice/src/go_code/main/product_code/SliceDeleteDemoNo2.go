package main

import "fmt"

func main() {
	/*声明一个切片*/
	var array = []int{1, 2, 3, 4, 5, 6}
	arrays := [...]int{1,2,3,4,5,6}

	fmt.Println(array[:0]) //相当于初始化了原切片
	fmt.Println(cap(array[:0]))

	/*使用append函数给切片赋值*/
	array = append(array[:0], array[1:]...)
	/*
	将目标切片的1号索引后的值添加到初始化的自身切片中
	保证了内存区域的一致性
	 */
	fmt.Println(arrays[0]) //通过数组索引找到对应的值打印出来
	fmt.Println("######")
	fmt.Println(array[2:3]) //打印出的是2，不是数组索引对应的值
	/*
	要注意区分数组的声明方式和切片的声明方式
	 */
	fmt.Println(copy(array, array[1:]))
	fmt.Println(array)
}