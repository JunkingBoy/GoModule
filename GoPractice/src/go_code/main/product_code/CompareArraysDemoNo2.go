package main

import "fmt"

var team [3]string

func main() {
	team[0] = "SpiderMan"
	team[1] = "HammerMan"
	team[2] = "SunWuKong"

	//使用切片的方法进行读取
	for i, j := range team {
		fmt.Printf("第%d位的值是%s\n", i, j)
	}
}