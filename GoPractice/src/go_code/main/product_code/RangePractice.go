package main

import "fmt"

func main() {
	/*声明一个切片*/
	slicePractice := []int{10,20,30,40}

	/*迭代值，显示key和value*/
	for i, j := range slicePractice {
		fmt.Printf("第%d号索引的值是%d\n", i, j)
	}

	/*显示值和地址*/
	for address, index := range slicePractice {
		fmt.Printf("Value: %d Value-Addr: %X ElemAddr: %X\n", index, &index, &slicePractice[address])
	}
}