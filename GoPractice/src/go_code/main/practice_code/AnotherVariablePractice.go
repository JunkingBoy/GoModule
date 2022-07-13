package main

import "fmt"

var (
	variableValueNo1 int
	variableValueNo2 int32
	variableValueNo3 float32
	variableValueNo4 int
	pointer *int
	pointerNo2 *int32
)

func main() {
	variableValueNo1 = 60
	fmt.Println("第一行 -a 变量类型为 = %T\n", variableValueNo1)
	fmt.Println("第二行 -b 变量类型 = %T\n", variableValueNo2)
	fmt.Println("第三行 -c 变量类型 = %T\n", variableValueNo3)

	pointer = &variableValueNo1 //'pointer'包含了变量variableValueNo1的地址
	/*如果在这改变指针的指向，那么结果为0*/
	pointer = &variableValueNo4
	fmt.Println("a 的值为 %d\n", variableValueNo1)
	fmt.Println("*pointer为 %d\n", *pointer)

	/*修改pointer的指向*/
	//初始化No4
	variableValueNo4 = 30
	pointer = &variableValueNo4
	fmt.Println("变量variableValueNo4的内存地址为:\n", pointer)

	variableValueNo2 = 30
	pointerNo2 = &variableValueNo2
	fmt.Println(*pointerNo2)
	/*
	注意：
	如果在指针变量前+*符号，那么打印出来的结果是这个变量指向的地址的值对象本身
	不+*符号，打印出来的结果是值对象的地址
	*/
}