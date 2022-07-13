package main

//单行注释
/*
多行注释
 */
import "fmt"

func main()  {
	fmt.Println("Lucifer" + "Jun")
	var number = 10
	var url = "https://home.cnblogs.com/u/JunkingBoy%s%d"
	var name = "/Lucifer"
	var answer = fmt.Sprintf(url, name, number)
	fmt.Println(answer)
	fmt.Println("Lucifer");fmt.Println("JunkingBoy")
	fmt.Println("Lucifer" + "Jun")
	var key, value bool
	fmt.Println(key, value)
	var arr [][]int
	fmt.Println(arr)
}