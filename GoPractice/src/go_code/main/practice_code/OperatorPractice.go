package main

import "fmt"

func main() {
	a := true
	b := false
	if a && b {
		fmt.Println("第一行条件为 ture\n")
	}
	if a || b {
		fmt.Println("第二行条件为 false\n")
	}

	/*修改a和b赋值*/
	a = b
	b = true
	if a && b {
		fmt.Println("第三行条件为 true\n")
	}else {
		fmt.Println("第三行条件为 false\n")
	}
	if !(a && b) {
		fmt.Println("第四行条件为 true\n")
	}
}