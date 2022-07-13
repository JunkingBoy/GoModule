package main

import (
	"fmt"
	"math"
)

func main() {
	/*打印数据类型的最大最小值*/
	fmt.Println("int8 range is:", math.MinInt8, math.MaxInt8)

	/*声明一个int32整数型*/
	var num int32 = 1047483647
	//格式化输出变量的16进制形式和10进制形式
	fmt.Printf("int32: 0x%x %d\n", num, num)

	/*数值转型为int16*/
	num2 := int16(num)
	fmt.Printf("int16: 0x%x %d\n", num2, num2)
	/*
	上下两次打印的结果上看发生了精度丢失
	 */

	/*声明一个float变量*/
	var num3 float32 = math.Pi //圆周率
	fmt.Println(int(num3)) //精度丢失，小数点后的数全部丢失
}