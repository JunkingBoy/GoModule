package main

import "fmt"

/*
创建一个创建匿名结构体并且打印的函数
 */
func printMes(mes *struct{ id int; data string }) {
	fmt.Printf("%T\n", mes)
}

func main() {
	//实例化匿名结构体
	msg := &struct {
		//定义成员部分
		id int
		data string
	}{
		//初始化值
		1024,
		"Hello",
	}

	//调用打印方法
	printMes(msg)
}