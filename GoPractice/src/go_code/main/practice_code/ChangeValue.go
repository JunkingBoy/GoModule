package main

import "fmt"

/*定义一个函数，实现传入两个整数，交换他们的值*/
func changeValue(arr, brr int) int {
	if arr == brr {
		return arr
	}else {
		//使用冒泡排序算法的方法让两个变量的值进行交换
		temp := arr
		arr = brr
		brr = temp
	}

	return brr
}

func main() {
	var valuePracticeNo1 int
	var valuePracticeNo2 int

	//获取输入值
	fmt.Println("First Value is :")
	fmt.Scanln(&valuePracticeNo1)
	fmt.Println("Second Value is :")
	fmt.Scanln(&valuePracticeNo2)

	fmt.Println("Before change value :", valuePracticeNo1, valuePracticeNo2)

	//调用方法
	fmt.Println("After change ValuePracticeNo2 value is :", changeValue(valuePracticeNo1,valuePracticeNo2))
}