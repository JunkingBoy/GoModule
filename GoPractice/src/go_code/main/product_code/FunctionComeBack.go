package main

import "fmt"

func comeBack() (a,b int) {
	return 1,2
}

func main() {
	fmt.Println(comeBack())
}