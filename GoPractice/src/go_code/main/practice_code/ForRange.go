package main

import "fmt"

func main() {
	/*声明一个映射*/
	practice := map[string]string{
		"name":"Jun",
		"littleName":"JunkingBoy",
	}

	for j,k := range practice {
		fmt.Println(j, k)
	}
}