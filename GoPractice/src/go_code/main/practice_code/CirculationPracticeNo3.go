package main

import "fmt"

var (
	intArray [4]int
)

func main() {
	//创建一个四个存储空间的数组
	intArray = [4]int{3,5,7,9}

	//使用range格式
	for key, number := range intArray{
		intArray[key] = number
		fmt.Println(intArray[key])
		fmt.Println(number)
	}

	numbers := [6]int{3,2,6,4,5,1}
	for i, j := range numbers{
		fmt.Println("第 %d 位的x的值为 %d\n", i, j)
	}

	for i := 0; i < len(numbers); i++ {
		for j := 0; j < len(numbers)-1-i; j++ {
			if numbers[j] > numbers[j+1] {
				//设置中间变量
				temp := numbers[j]
				numbers[j] = numbers[j+1]
				numbers[j+1] = temp
			}
		}
	}
	for i, j := range numbers{
		fmt.Println("索引",i, "对应的值为", j)
	}

	num := 10
	Loop:
		for num < 20 {
			if num == 15 {
				//自增，回到标签处
				num += 1
				goto Loop
			}
			fmt.Println(num)
			num++
		}
}