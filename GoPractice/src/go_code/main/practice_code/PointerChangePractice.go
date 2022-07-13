package main

import "fmt"

/*写一个函数，交换传入的指针变量的值*/
func changePointerValue(arr *int, brr *int) {
	//指针之间进行值的交换。
	//使用中间变量的方法存储值
	temp := *arr
	*arr = *brr
	*brr = temp
}

func main() {
	//设置两个变量，指定变量的值。通过交换指针的指向达到改变两个变量的值
	a := 200
	b := 100

	fmt.Printf("指针交换前a的值:\n", a)
	fmt.Println("\n")
	fmt.Printf("指针交换前b的值:\n", b)

	/* 调用 swap() 函数
	 * &a 指向 a 指针，a 变量的地址
	 * &b 指向 b 指针，b 变量的地址
	 */
	changePointerValue(&a, &b)

	fmt.Println("\n")
	fmt.Println("指针交换后a的值:", a)
	fmt.Println("指针交换后b的值:", b)
}