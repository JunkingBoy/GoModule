package main

import "fmt"

func main() {
	/*处理第几行*/
	for y := 1; y <= 9; y++ {
		/*每一行生成多少列*/
		for x := 1; x <= y; x++ {
			//打印结果
			fmt.Printf("%d*%d=%d\t", y, x, y*x)
		}
		//生成回车
		fmt.Println()
	}
}