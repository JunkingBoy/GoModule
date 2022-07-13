package cls1

import (
    "GoPracticeClsFactory/base"
    "fmt"
)

/*
定义一个结构体
实现base包下的class接口
定义初始化的init()函数
 */

/* 定义一个结构体1 */
type Class1 struct {
}

/* 实现factory下的class接口 */
func (c *Class1) Do() {
    fmt.Println("HelloWorld!")
}

/* 定义init()初始化函数 */
func init() {
    base.Register("Class1", func() base.Class{
        return new(Class1)
    })
}
