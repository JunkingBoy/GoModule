package main

import "fmt"

func main() {
	/*声明一个map*/
	mapValue := map[string]int{}

	/*为其赋值*/
	mapValue["age"] = 22
	mapValue["long"] = 18
	mapValue["howlong"] = 16

	/*打印map*/
	fmt.Println(mapValue)

	/*试用delete函数删除掉一个key*/
	delete(mapValue, "howlong")

	fmt.Println(mapValue)
}