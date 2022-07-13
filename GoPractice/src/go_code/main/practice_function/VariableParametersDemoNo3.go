package main

import (
	"bytes"
	"fmt"
)

func jointString(slist ...string) string {
	//定义变量--->这个变量相当于字符串拼接完成后的值--->这是一个字节数组
	var b bytes.Buffer

	//循环获取形参数组里面的内容
	for _, v := range slist {
		//将遍历出的字符串连续写入字节数组
		b.WriteString(v)
	}

	return b.String()
}

func main() {
	//连续输入字符将它们拼接成一个字符串
	fmt.Println(jointString("Jun", "is", "a", "handsome", "boy!"))
}