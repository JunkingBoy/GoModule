package main

import "fmt"

func main() {
	/*声明一个多维数组*/
	var multiarray [][]int

	//给多维数组赋值
	multiarray = [][]int{{10}, {100,200}}

	fmt.Println(multiarray)

	fmt.Println(multiarray[0][0])

	fmt.Println(multiarray[1][0])
	fmt.Println(multiarray[1][1])

	fmt.Println(cap(multiarray)) //容量为2
	fmt.Println(len(multiarray[0])) //多维数组的零号索引处的长度为1
	fmt.Println(len(multiarray[1])) //多维数组的一号索引处的长度为2
}