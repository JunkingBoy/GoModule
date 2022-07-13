package main

import "fmt"

var globalVariable int //这是一个全局变量

func main() {
	c := 3 //这是一个局部变量
	/*
	声明了全局变量但不使用
	 */
	fmt.Println(c)
}