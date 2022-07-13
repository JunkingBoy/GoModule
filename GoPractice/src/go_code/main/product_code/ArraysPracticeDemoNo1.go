package main

import "fmt"

func main() {
	var arrays [3]int

	for i := 0; i < len(arrays); i++ {
		arrays[i] = i+1
	}

	//打印索引+元素
	for j, k := range arrays {
		fmt.Printf("第%d号索引的值是%d;\n", j, k)
	}

	//仅打印值--->使用_
	for _, l := range arrays {
		fmt.Printf("依次的值是%d;\n", l)
	}
}