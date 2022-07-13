package main

import "fmt"

func main() {
	/*声明三个Map*/
	var mapList map [string] int
	var mapAssigned map [string] int
	mapCreated := make(map [string] float32)

	/*往mapList中添加元素*/
	mapList = map [string] int{"long": 18, "age": 22}

	//因为mapAssigned的key类型和value类型和mapList一致，所以可以直接赋值
	mapAssigned = mapList

	//另一种对map赋值的方式
	mapCreated["long"] = 3.1415926
	mapCreated["age"] = 18
	mapCreated["big"] = 19

	//格式化输出值
	fmt.Printf("Map List at \"long\" is :%d\n", mapList["long"])
	fmt.Printf("Map List at \"age\" is :%d\n", mapList["age"])
	fmt.Printf("Map Assigned \"age\" is :%d\n", mapAssigned["age"])
	fmt.Printf("Map Created \"big\" is :%f\n", mapCreated["big"])
	fmt.Printf("Map Created \"big\" is :%f\n", mapCreated["long"])
}