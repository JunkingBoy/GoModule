package main

import "fmt"

func parameters(args ...interface{}) {
	for _, v := range args {
		switch v.(type) {
		case string:
			fmt.Printf("Value is:%s\n", v)
		case int:
			fmt.Printf("Value is:%d\n", v)
		case float32:
			fmt.Printf("Value is:%f\n", v)
		case float64:
			fmt.Printf("Value is:%f", v)
		default:
			fmt.Println("Unknow type!!!")
		}
	}
}

func main() {
	parameters(1,"JunkingBoy",3.1415)
}