package main

import "fmt"

const (
	valueNo1 = iota //(adj 极小的)第一个iota=0
	valueNo2
	valueNo3
	valueNo4 = "Jun" //这是一个独立的值，但是iota还是在计数--->这里的iota += 1(实际在这里iota=3)
	valueNo5 //这里如果不显式声明valueNo5的值，那么它将和valueNo4指向一个对象
	valueNo6 = 100
	valueNo7
	valueNo8 = iota
	valueNo9
)

func main() {
	fmt.Println(valueNo1,valueNo2,valueNo3,valueNo4,valueNo5,valueNo6,valueNo7,valueNo8,valueNo9)
}