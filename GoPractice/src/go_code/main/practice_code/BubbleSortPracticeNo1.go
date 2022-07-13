package main

import "fmt"

var (
	temp int
)

func bubbleSort(arr []int) {
	if len(arr)==0 {
		fmt.Println("输入错误!!!")
	}else if len(arr)==1 {
		fmt.Println(arr[0])
	}else {
		for i:=0; i<len(arr); i++ {
			for j:=0; j<len(arr)-1-i; j++ {
				if arr[j]>arr[j+1] {
					temp = arr[j]
					arr[j] = arr[j+1]
					arr[j+1] = temp
				}
			}
		}
	}
	for i, num := range arr{
		fmt.Println("The sorting result is:" , i, num)
	}
}

func main() {
	arr := []int{6,3,2,7,9}
	bubbleSort(arr)
}