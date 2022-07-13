package main

import (
	"fmt"
)

func visit(list []int, f func(int)) {
	/*循环取出切片的value放入切片*/
	for _, v := range list {
		f(v)
	}
}

func main() {
	/*使用匿名函数打印切片内容*/
	visit([]int{7,8,9,1,2,3}, func(i int) {
		fmt.Println(i)
	})
}