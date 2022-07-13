package main

import "fmt"

func main() {
	fmt.Println("Begin")

	//使用defer语句类似于栈指针
	defer fmt.Println(1)

	defer fmt.Println(2)

	defer fmt.Println(3) //此时这个函数是最后入栈的，相当于位于栈顶。先出

	fmt.Println("End")
}