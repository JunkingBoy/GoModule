package main

import (
	"bytes"
	"fmt"
	"strings"
)

func main() {
	subString := "同学起立，老师上课。"

	//截取出来的字符串
	subStringNo1 := strings.Index(subString, "，") //这只是截取出来了字符串里

	subStringNo2 := strings.Index(subString[subStringNo1:], "人类")

	fmt.Println(subStringNo1, subStringNo2, subString[subStringNo1+subStringNo2:])

	hammer := "锤哥"

	//声明字符缓冲变量
	stringBuilder := bytes.Buffer{}

	//把字符串写进来
	stringBuilder.WriteString(hammer) //字符串转变成了二进制形式

	fmt.Println(stringBuilder)

	//将缓冲以字符串的形式输出
	fmt.Println(stringBuilder.String())
}