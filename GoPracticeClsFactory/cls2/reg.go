package cls2

import (
    "GoPracticeClsFactory/base"
    "fmt"
)

/*
定义结构体
实现factory文件下的接口
定义init()方法
 */

/* 定义结构体 */
type Class2 struct {
}

func (c *Class2) Do() {
    fmt.Println("Class2")
}

func init() {
    // 启动注册器自动注册class2工厂--->调用base包下定义的函数
    base.Register("Class2", func() base.Class {
        return new(Class2)
    })
}
