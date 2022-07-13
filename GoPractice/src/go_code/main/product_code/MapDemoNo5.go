package main

import (
	"fmt"
	"sort"
)

func main() {

	/*声明一个mapPair*/
	mapValue := map [string] int{}

	/*添加值*/
	mapValue["long"] = 25
	mapValue["width"] = 22
	mapValue["height"] = 20
	mapValue["weight"] = 18

	/*遍历key和value*/
	for i, j := range mapValue {
		fmt.Printf("key是%s,值是%d:\n", i, j)
	}

	/*只遍历value*/
	for _, k := range mapValue {
		fmt.Printf("全部的值是%d\n", k)
	}

	/*只遍历key*/
	for l := range mapValue {
		fmt.Printf("全部的key是%s\n", l)
	}

	/*声明一个切片*/
	var srcSlice []string

	//复制key到srcSlice当中
	for h := range mapValue {
		srcSlice = append(srcSlice, h)
	}

	//对切片排序
	sort.Strings(srcSlice)

	fmt.Println(srcSlice)
}