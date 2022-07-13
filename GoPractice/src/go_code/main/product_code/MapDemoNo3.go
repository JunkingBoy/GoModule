package main

import "fmt"

func main() {
	/*声明一个Map映射*/
	mapValue := make(map [string] int)

	mapValue2 := make(map [int] *[]int)

	intValue := []int{1,2,3,4,5,6}

	mapValue2 = map [int] *[]int{
		1: &intValue,
	}
	mapValue2[2] = &intValue
	/*Map有两种赋值方式*/
	/*
	1、像mapValue = map [string] int{"age": 22}再value后接"{}"里面写key和value
	2、像mapValue2[2] = &intValue这样，将value赋值给左边的key位--->"[]"里面写key
	 */

	mapValue = map [string] int{"age": 22}

	fmt.Println(mapValue)
	fmt.Println(mapValue2)
}