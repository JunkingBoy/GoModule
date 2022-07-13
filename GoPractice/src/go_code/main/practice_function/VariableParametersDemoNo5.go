package main

import "fmt"

/*
第一个函数
 */
func rawPrint(slist ...interface{}) {
	//遍历切片
	for _, v := range slist {
		//打印参数
		fmt.Println(v)
	}
}

/*
第二个函数，在内部调用第一个函数
 */
func print(slist ...interface{}) {
	//将传入的形参元素完整的传递给第一个函数
	rawPrint("fmt", slist)
}

func main() {
	//外部调用第二个函数
	print(1, "JunkingBoy", 2)
}