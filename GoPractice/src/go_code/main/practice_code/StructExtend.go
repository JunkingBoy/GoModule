package main

import "fmt"

/*
定义一个飞行结构体，里面有属于该结构体的飞行方法
 */
type Flying struct {

}

/*
飞行方法
 */
func (f *Flying) Fly() {
    fmt.Println("Can Fly!")
}

/*
移动结构体
 */
type Move struct {

}

/*
行走方法
 */
func (m *Move) Walk() {
    fmt.Println("Walking on the road!")
}

/*
鸟类，能走能飞，内嵌两个结构体
 */
type Bird struct {
    Flying
    Move
}

/*
人类，能走不能飞，内嵌一个结构体
 */
type Human struct {
    Move
}

func main() {
    //实例化结构体
    br := new(Bird)
    br.Fly()
    br.Walk()
    fmt.Println("#####################")
    hu := new(Human)
    hu.Walk()
}