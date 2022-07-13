package main

import "fmt"

/*
定义一个飞行结构体--->和Java当中定义类不一样，是一种存在于类之上的抽象的东西
 */
type Flying struct {
	
}

/*
定义这个飞行结构体能够拥有的方法--->类似于Java中类的方法
由于结构体当中的特性，结构体当中只能够定义成员不能够定义函数
 */
func (f *Flying) Fly() {
	fmt.Println("can fly")
}

/*
定义一个可行走的结构体--->行走类
 */
type Walking struct {
	
}

/*
定义这个行走结构体当中的方法
 */
func (w *Walking) Walk() {
	fmt.Println("can walk")
}

/*
定义人类，人类能够行走不能飞--->到这里定义的人类的结构体才是类似于Java当中的类
内嵌的结构体相当于类当中的方法。只是Go当中使用了另外一种展现了面向对象的思想
 */
type Human struct {
	Walking
}

/*
定义鸟类
 */
type Bird struct {
	Flying
	Walking
}

func main() {
	//实例化人类
	hu := new(Human)
	fmt.Println("Human is:")
	hu.Walk() //访问到内嵌结构体当中的方法
	fmt.Println("##########")
	//实例化鸟类
	br := new(Bird)
	fmt.Println("Bird is:")
	br.Fly()
	br.Walk()
}