package main

import "fmt"

var (
	chooseArray []int
)

func main() {
	chooseArray = []int{3,6,7,1,2,5}
	chooseSort(chooseArray)
}

func chooseSort(arr []int) {
	for i := 0; i < len(arr); i++ {
		//找到最小值--->这里的i是假设下表i为最小。实际交换的一直是数组下标
		//之所以这样赋值是保证每次循环进来比较的都会从排好序的下一位开始比较
		min := i
		//从i的下一位开始循环和i进行比较
		for j := i+1; j < len(arr); j++ {
			if arr[j] < arr[min] {
				//下标交换
				min = j
			}
		}
		//判断如果最小值!=i了，将最小值与i位置进行数据替换，将最小值放到数组前面来
		if min != i {
			//使用冒泡排序的交换值得方法
			temp := arr[i]
			arr[i] = arr[min]
			arr[min] = temp
		}
	}
	//打印出排序好得值
	for i, j := range arr {
		fmt.Println("索引", i, "得值为:", j)
	}
}