package main

import "fmt"

/*
可以称为一个外函数
 */
func Accumulate(value int) func() int {
	//返回一个内部函数，内部函数引用外部函数的形参value形成闭包
	return func() int {
		//引用value
		value++
		//返回value
		return value
	}
}

func main() {
	//创建一个累加器，初始值为1
	accumulate1 := Accumulate(1)

	//打印变量
	fmt.Println(accumulate1())

	//打印变量的函数地址
	fmt.Printf("Variable address is:%p\n", &accumulate1)

	//在实例化一个外函数
	accumulate2 := Accumulate(10)

	//打印变量
	fmt.Println(accumulate2())

	//打印变量地址
	fmt.Printf("Variable address No2 is:%p\n", &accumulate2)
	fmt.Printf("Variable address is:%p\n", &accumulate1)
}