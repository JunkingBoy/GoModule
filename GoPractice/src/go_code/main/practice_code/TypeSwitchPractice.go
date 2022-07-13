package main

import (
	"fmt"
	"go/types"
)

func main() {
	var inter interface{} = true

	//使用变量去代替接口当中的值并且判断类型
	switch i := inter.(type) {
	case types.Nil:
		fmt.Println("x的类型是:", i)
	case int:
		fmt.Println("x是int类型")
	case float64:
		fmt.Println("x是float64类型")
	case func(int2 int):
		fmt.Println("x是func(int)类型")
	case bool, string:
		fmt.Println("x是bool或string类型")
	default:
		fmt.Println("未知类型")
	}
}