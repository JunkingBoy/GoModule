package main

var (
	n int
)

func insertSort(arr []int) {
	n = len(arr)

	//第一层循环寻找位置进行比较-->不断循环到数组长度最后一位索引
	for i := 0; i < n; i++ {
		//内层循环开始寻找插入的值的位置--->依次和当前位置之前的值进行比较
		for j := i; j > 0; j-- {
			if (arr[j] - arr[j-1]) < 0 {
				//调用换位置方法
				changePlace(arr, j, j-1)
			}
		}
	}
}

func changePlace(brr []int, k int, l int) {
	temp := brr[l]
	brr[l] = brr[k]
	brr[k] = temp
}