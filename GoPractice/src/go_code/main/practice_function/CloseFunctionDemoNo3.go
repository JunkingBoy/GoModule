package main

import "fmt"

/*
定义一个玩家函数，里面定义内部函数通过闭包的方式返回变量值
@date 2021/09/06
@author Lucifer
 */
func playerGame(name string) func() (string, int) {
	//定义hp
	hp := 150
	//通过闭包返回局部变量和外部函数的引用
	return func() (string, int) {
		return name, hp
		/*
		这里面的name是对外部函数形参的引用，hp是对外部函数的局部变量的引用
		因此形成了闭包
		 */
	}
}

func main() {
	//将外部函数定义成变量
	player1 := playerGame("God")
	//定义闭包中的两个返回值属性
	name, hp := player1()
	//打印这两个变量
	fmt.Println(name,hp)
}