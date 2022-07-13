package main

import "fmt"

var x, y int
var (
	a int
	b bool
	c string
	d float64
	ptr *int
	ptrPointer **int
	channel chan int
	chaPointer **chan int
)

var e, f = 1, "JunkingBoy"

func main() {
	g, h := "Hello", "World"
	fmt.Println(a,b,c,d,e,f,x,y,g,h)
	arrays := []int{1,2,3,4,5,6}
	fmt.Println(arrays)
	ptr = &a
	ptrPointer = &ptr //注意这里不能直接用赋值符号，应该使用的是引用符号。既引用*ptr指针的地址
	fmt.Println("指向指针的指针变量值为:", **ptrPointer)
	fmt.Println("它的地址为:", ptrPointer)
}