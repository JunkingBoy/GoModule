package main

import (
	"fmt"
	"regexp"
	"strconv"
)

/*
Go语言中编译正则表达式使用的是regexp包下的函数：
1、Compile
2、MustCompile
 */
func MustCompile(str string) *regexp.Regexp {
	//将形参传入正则包下的方法进行转换
	regexp, error := regexp.Compile(str)
	if error != nil {
		//手动宕机
		panic(`regexp: Compile(` + strconv.Quote(str) + `): ` + error.Error())
	}
	//error为空，返回正则对象
	return regexp
}

func main() {
	fmt.Println(MustCompile("JunkingBoy"))
}