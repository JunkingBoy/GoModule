package main

import (
	"fmt"
)

func main() {

}

func chooseSortNo2() (arr []int) {
	var numbers int
	chooArray := [5]int{}
	fmt.Println("You can input five numbers\n")
	for i := 0; i < 5; i++ {
		fmt.Println("Input Array:", i, "Value:")
		fmt.Scanln(&numbers)
		chooArray[i] = numbers
	}

	//读取排序前的索引
	for i, j := range chooArray {
		fmt.Println("Before Sort No.",i ,"Value is :", j)
	}

	//进行选择排序--->将输入的内容自动排序
	for i := 0; i < len(chooArray); i++ {
		//设置最小值的索引
		minValue := i
		//内层循环与最小值进行比较，并且确保每一次获取的都是排序号以后的下一位索引
		for j := i + 1; j < len(chooArray); j++ {
			//和最小值进行比较
			if chooArray[j] < chooArray[minValue] {
				//把最小值的索引直接赋值给当前的最小值索引
				minValue = j
			}
		}
		//如果最小值的下标不等于i了，那么将最小值与i位置的数据替换，即将最小值放到数组前面来，然后循环整个操作。
		if minValue != i {
			//使用冒泡排序的方法进行值得交换排序
			temp := chooArray[i]
			chooArray[i] = chooArray[minValue]
			chooArray[minValue] = temp
		}
	}

	//打印分隔符
	fmt.Println("==========")

	//索引后的值为
	for i, j := range chooArray {
		fmt.Println("After Sort No.", i, "Value is :", j)
	}
	return arr
}