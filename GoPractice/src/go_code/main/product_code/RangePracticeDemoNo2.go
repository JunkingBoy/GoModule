package main

import "fmt"

func main() {
	/*声明一个切片*/
	slicePractice := []int{10,20,30,40}

	for index := range slicePractice {
		fmt.Printf("Index value is %d\n", index)
		fmt.Printf("引用地址是%x\n", &index)
	}
}