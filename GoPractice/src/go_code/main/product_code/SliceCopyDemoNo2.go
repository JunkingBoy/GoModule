package main

import "fmt"

func main() {
	//设置元素数量
	const elementNumbers = 1000

	//使用make()方法预分配足够多的切片
	srcData := make([]int, elementNumbers)

	//将数组赋值
	for i := 0; i < elementNumbers; i++ {
		srcData[i] = i+1
	}

	//打印出切片来看看
	//for i, j := range srcData {
	//	fmt.Printf("第%d号索引的值是%d\n", i, j)
	//}

	//引用切片数据
	refData := srcData

	//使用make预分配切片
	copyData := make([]int, elementNumbers)

	//使用copy内置函数复制src给copy
	copy(copyData, srcData)

	//修改第一个元素
	srcData[0] = 999

	srcData[1] = 888

	//打印引用切片的第一个元素
	fmt.Println(refData[1])

	copyData[0] = 777

	//打印第一个和最后一个元素
	fmt.Println(copyData[0], copyData[elementNumbers-1])

	srcData[4] = 444

	srcData[5] = 666

	// 复制原始数据从4到6(不包含)
	copy(copyData, srcData[4:6])

	for i := 0; i < 5; i++ {
		fmt.Printf("%d", copyData[i])
	}
	/*
	引用的话源切片和目的切片的内存地址指向的连续内存区域是一样的
	改变源切片的元素目的切片也会对应的改变
	复制的话原切片和目的切片的内存地址指向的连续内存区域是不一样的
	复制是新开辟出一块区域
	 */
}