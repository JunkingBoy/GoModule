package main

import "fmt"

func main() {
	var arraysNo2 [...]int

	//循环赋值
	for i := 0; i < 10; i++ {
		arraysNo2[i] = i+1
	}

	//循环输出
	for j, k := range arraysNo2 {
		fmt.Printf("第%d号索引的值是%d;\n", j, k)
	}
}