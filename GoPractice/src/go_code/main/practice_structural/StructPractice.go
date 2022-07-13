package main

/*
定义一个point结构体
 */
type Point struct {
	X float64
	Y float64
	Z float64
}

func main() {
	//利用变量的形式把point结构体实例化出来
	var p Point
	p.X = 3.14
	p.Y = 6.28
	p.Z = 9.72
}