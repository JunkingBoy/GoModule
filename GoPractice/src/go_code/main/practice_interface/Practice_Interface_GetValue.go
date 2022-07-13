package main

import "fmt"

func main() {
    // 声明一个int类型的变量
    var a int = 1

    // 声明一个接口，将i赋值给接口
    var i interface{} = a

    // 声明变量b然后直接赋值--->编译不通过
    /*
    将 a 的值赋值给 i 时，虽然 i 在赋值完成后的内部值为 int，但 i 还是一个 interface{} 类型的变量
     */
    //var b int = i

    // 声明变量b通过类型断言赋值--->编译通过
    var b int = i.(int)
    fmt.Println(b)
}
