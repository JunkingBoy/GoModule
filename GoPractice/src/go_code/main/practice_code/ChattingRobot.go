package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

/*使用go语言的基本语法实现一个聊天机器人*/
/*
主要运用到的知识：
1、io中的输入捕获
2、string数据类型
3、switch...case语句
*/

func getName() {
	inputReader := bufio.NewReader(os.Stdin) //准备从标准输入读取数据
	fmt.Println("Tell me who are you:")
	input, err := inputReader.ReadString('\n') //读取数据直到遇到\n
	//判断并抛出异常
	if err != nil {
		fmt.Println("An error occurred!!!")
		os.Exit(500)
	}else {
		//用切片操作删除最后的 \n
		name := input[:len(input)-1]
		fmt.Println("Hello,", name, "Are you a boy or a girl?\n")
	}

	//循环进行输入
	for {
		input, err = inputReader.ReadString('\n')
		if err != nil {
			fmt.Printf("An error occurred:%s\n",err)
			continue
		}
		input = input[:len(input)-1]
		//全部转成小写
		input = strings.ToLower(input)
		if input == "" {
			continue
		}else if input == "boy" {
			fmt.Println("Good boy!")
		}else if input == "girl" {
			fmt.Println("Good girl!")
		}else if input == "bye" {
			fmt.Println("Good bye my friend!")
			os.Exit(0)
		}else {
			fmt.Println("Sorry,I don't know what are you say!")
		}
	}
}

func main() {
	getNameNo2()
}

func getNameNo2() {
	inputReader := bufio.NewReader(os.Stdin)  //准备从标准输入读取数据
	fmt.Println("Please input your name:")
	input,err := inputReader.ReadString('\n')  //读取数据直到碰到 \n为止
	if err != nil{
		fmt.Printf("An error occurred:%s")
		os.Exit(1)  //异常退出
	}else {
		//用切片操作删除最后的 \n
		name := input[:len(input)-1]
		fmt.Printf("Hello,%s! What can I do for you?\n",name)
	}
	for{
		input,err = inputReader.ReadString('\n')
		if err != nil{
			fmt.Printf("An error occurred:%s\n",err)
			continue
		}
		input = input[:len(input)-1]
		//全部转换为小写
		input = strings.ToLower(input)
		switch input {
		case "":
			continue
		case "nothing","bye":
			fmt.Println("Bye!")
			//正常退出
			os.Exit(0)
		default:
			fmt.Println("Sorry,I didn't catch you.")
		}
	}
}