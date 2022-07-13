package main

import "fmt"

func loop()  {
	for i := 0; i < 10; i++ {
		fmt.Println("%d", i)
	}
}

func main() {
	go loop() //启动一个goroutine
	loop()
}
