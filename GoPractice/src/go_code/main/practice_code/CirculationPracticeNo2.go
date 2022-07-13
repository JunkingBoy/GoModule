package main

import "fmt"

func main() {
	i := 0
	for {
		i++
		fmt.Println(i)
		if i == 65535 {
			fmt.Println("收工!!!")
			break
		}
	}
}