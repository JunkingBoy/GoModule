package main

import "fmt"

/*
结构体A
 */
type A struct {
    a int
}

/*
结构体B
 */
type B struct {
    a int
}

/*
内嵌A和B的结构体C
 */
type C struct {
    A
    B
}

func main() {
    //实例化结构体C赋值给内嵌结构体A的属性a
    c := &C{}
    c.A.a = 1
    c.B.a = 1
    //c.a = 1
    fmt.Println(c)
}