package main

import (
	"bytes"
	"fmt"
)

func printTypeValue(slist ...interface{}) string {
	//定义字节数组
	var b bytes.Buffer
	//定义一个描述类型的字符串
	var valueType string

	//循环获取interface{}当中的值
	for _, s := range slist {
		//定义一个写入的值的变量
		str := fmt.Sprintf("%v", s) //注意这个方法是Sprintf不是Printf
		//判断值的类型
		switch s.(type) {
		case bool:
			valueType = "bool"
		case int:
			valueType = "int"
		case string:
			valueType = "string"
		default:
			fmt.Println("Unknow value type!")
		}

		//分别写入字符串前缀、值、类型
		b.WriteString("value is:")
		b.WriteString(str + "\t")
		b.WriteString("type:")
		b.WriteString(valueType)
		b.WriteString("\n")
	}

	return b.String()
}

func main() {
	fmt.Println(printTypeValue(100, "JunkingBoy", false))
}