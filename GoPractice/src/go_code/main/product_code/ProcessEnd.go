package main

import (
	"flag"
	"fmt"
)

var skillParam = flag.String("skill", "", "skill to perform")

func main() {
	/*解析命令行参数*/
	flag.Parsed()

	//利用匿名函数赋值给skill变量
	var skill = map[string]func(){
		"fire": func() {
			fmt.Println("chicken fire")
		},
		"run": func() {
			fmt.Println("soldier run")
		},
		"fly": func() {
			fmt.Println("angel fly")
		},
	}

	//var f func() = skill[*skillParam]
	//
	//var ok bool
	//
	//if ok {
	//	f()
	//}else {
	//	fmt.Println("skill not found")
	//}

	if f, ok := skill[*skillParam]; ok {
		f()
	}else {
		fmt.Println("skill not found")
	}
}