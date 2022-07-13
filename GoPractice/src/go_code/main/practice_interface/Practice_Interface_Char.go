package main

import "fmt"

/*
定义一个接口
 */
type Student interface {
    ShowInfor() string
}

/*
定义一个类型
 */
type stu struct {
    name string
    age int
}

/*
该类型实现ShowInfor函数
 */
func (s *stu) ShowInfor() string {
    return "Hi!"
}

/*
定义一个函数，返回该接口类型
 */
func getStudents() Student {
    // 声明stu类型的指针变量--->带有类型的nil（类型为：*stu）
    var s *stu = nil
    if nil == s {
        return nil
    }
    return s
}

func main() {
    fmt.Println("==========显式将nil赋值给接口==========")
    var ss Student
    ss = nil
    if nil == ss {
        fmt.Println("收到的信息是：nil")
    }else {
        fmt.Println("收到的信息是：not nil")
    }

    fmt.Println("==========带有类型的nil赋值给接口==========")
    // 调用getStudents方法
    if nil == getStudents() {
        fmt.Println("收到的信息是：nil")
    }else {
        fmt.Println("收到的信息是：not nil")
    }
}
