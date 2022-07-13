package main

import "fmt"

func getValue(args ...int)  {
	for _, v := range args {
		fmt.Printf("Value is:%d\n", v)
	}
	/*
	这里的形参带了...会自动将数据类型转为数组类型使用切片的方式进行接收和打印
	 */
}

/*
如果没有...的语法糖那么在实现的时候需要这样写
 */
func getValue2(args []int) {
	for _, v := range args {
		fmt.Printf("Value is:%d\n", v)
	}
}

func main() {
	getValue(1,2,3,4,5)
	fmt.Println("##########")
	//调用的时候传入的参数就必须是数组切片
	getValue2([]int{5,4,3,2,1})
}