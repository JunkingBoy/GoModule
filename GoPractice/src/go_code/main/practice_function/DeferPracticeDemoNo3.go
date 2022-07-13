package main

import (
	"fmt"
	"os"
)

func fileSize(fileName string) int64 {
	//根据文件名打开文件, 返回文件句柄和错误
	f, err := os.Open(fileName)

	if err != nil {
		return 0
	}

	//获取文件状态
	info, err := f.Stat()
	if err != nil {
		//关闭资源
		f.Close()
		return 0
	}

	//取文件大小
	size := info.Size()

	//关闭资源
	f.Close()

	//返回名称和大小
	return size
}

/*
使用defer简化上面的函数。这里要注意：
1、defer是放在函数中最后执行的语句，相当于位于栈底
2、defer的声明式显式的
 */
func fileSize2(fileName string) int64 {
	//打开资源
	f, err:=os.Open(fileName)

	if err != nil {
		//返回0
		return 0
	}

	//正常打开资源接下来就是获取资源
	//声明一个遇到异常关闭函数的释放资源的方法
	defer f.Close()

	//获取文件状态
	info, err:=f.Stat()
	if err != nil {
		//此时会结束函数，结束前会运行defer标记的内容，所以这里无需声明f.close()
		return 0
	}

	//获取文件大小
	size := info.Size()

	//返回size
	return size
}

func main() {
	fmt.Println(fileSize("APITest.jar"))
}