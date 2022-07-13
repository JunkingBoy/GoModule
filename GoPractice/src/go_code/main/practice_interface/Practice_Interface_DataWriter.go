package main

import "fmt"

// 定义数据写入器
type DataWriter interface {
    // 定义接口中的方法
    WriterData(data interface{}) error
}

// 定义结构体。用于实现DataWriter接口。类似于java中的类
type file struct {

}

// 单独定义接口中方法的内容(实现接口中定义的方法)
/*
因为没有类似implements这样的关键字，所以在方法名前加入该结构体的指针对象
表示使用该结构体实现接口WriterData
file 的 WriteData() 方法使用指针接收器。输入一个 interface{} 类型的 data，返回 error。
 */
func (d *file) WriterData(data interface{}) error {
    // 写入数据
    fmt.Println("WriterData:", data)
    // 返回值
    return nil
}

func main() {
    // 实例化file
    f := new(file)

    // 声明一个接口变量
    var writer DataWriter

    // 将实现了接口的结构体对象赋值给接口
    /*
    将 *file 类型的 f 赋值给 DataWriter 接口的 writer，虽然两个变量类型不一致。但是 writer 是一个接口，且 f 已经完全实现了 DataWriter() 的所有方法，因此赋值是成功的
     */
    writer = f

    // 使用writer接口变量调用方法进行数据写入
    writer.WriterData("data")
}
