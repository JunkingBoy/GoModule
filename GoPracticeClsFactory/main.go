package main

/*
调用前面的三个目录下定义文件下的函数
 */
import (
    "GoPracticeClsFactory/base"
)

func main() {
    // 创建c1实例
    c1:= base.Create("Class1")
    // 调用接口下的函数
    c1.Do()

    // 创建c2实例
    c2 := base.Create("Class2")
    c2.Do()
}
