package main

import "fmt"

func main() {
	number := 8
	name := "Jun"

	fmt.Printf("带噶后，我今年%d岁了，我叫%s", number, name)

	pi := 3.14159
	// 按数值本身的格式输出
	variant := fmt.Sprintf("%v %v %v", "月球基地", pi, true)
	fmt.Println(variant)
	// 匿名结构体声明, 并赋予初值
	profile := &struct {
		Name string
		HP   int
	}{
		Name: "rat",
		HP:   150,
	}
	fmt.Printf("使用'%%+v' %+v\n", profile)
	fmt.Printf("使用'%%#v' %#v\n", profile)
	fmt.Printf("使用'%%T' %T\n", profile)

	string := "JunkingBoy"

	fmt.Println(len(string)-1)
	/*
	C语言中, 使用%d代表整型参数
	 */
}