package main

import (
	"flag"
	"fmt"
)

var mode = flag.String("mode", "", "process mode")
/*
	通过 flag.String，定义一个 mode 变量，这个变量的类型是 *string。后面 3 个参数分别如下：
		参数名称：在命令行输入参数时，使用这个名称。
		参数值的默认值：与 flag 所使用的函数创建变量类型对应，String 对应字符串、Int 对应整型、Bool 对应布尔型等。
		参数说明：使用 -help 时，会出现在说明中。
 */

func main() {
	//解析命令行参数
	flag.Parse() //解析命令行参数，并将结果写入到变量 mode 中。

	//输出命令行参数
	fmt.Println(*mode) //打印 mode 指针所指向的变量。
}
/*
之前已经使用 flag.String 注册了一个名为 mode 的命令行参数，flag 底层知道怎么解析命令行，并且将值赋给 mode*string 指针
在 Parse 调用完毕后，无须从 flag 获取值，而是通过自己注册的这个 mode 指针获取到最终的值。
 */