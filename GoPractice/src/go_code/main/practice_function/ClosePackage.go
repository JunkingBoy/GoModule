package main

/*
实现一个闭包函数，打印形参的值
 */
func close(arr int) func() {
	return func() {
		print(arr)
	}
}

func main() {
	//设置函数变量
	f := close(100)
	//调用函数
	f()
}