package main

import "fmt"

func main() {
	/*声明一个预定义长度的map*/
	mapValue := make(map [string] float32, 100)
	mapValue = map [string] float32{
		"C0": 16.35, "D0": 18.35, "E0": 20.60, "F0": 21.83,
		"G0": 24.50, "A0": 27.50, "B0": 30.87, "A4": 440,
	}

	fmt.Println(mapValue)
}