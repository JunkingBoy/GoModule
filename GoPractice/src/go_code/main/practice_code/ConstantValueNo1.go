package main

import (
	"fmt"
	"unsafe"
)

/*
go语言之常量的练习
1、关键字
2、声明方法
 */
const VALUE = 1
const VALUENO2 string = "Jun"
const (
	LENGTH int = 10
	WIDTH int = 5
	A = "Lucifer"
	B = len(A)
	C = unsafe.Sizeof(A)
)

const (
	h = iota
	i
	j
)

var (
	area int
)

func main() {
	const a, b, c = 1,false,"Jun"
	const e,f,g int = 1,2,3
	fmt.Println(LENGTH,WIDTH,area,a,b,c,e,f,g)
	fmt.Println(A,B,C)
	fmt.Println(h,i,j)
}