package main

import "fmt"

func bubbleSortNo2(brr []int) {
	for i := 0; i < (len(brr)-1); i++ {
		for j := 0; j < (len(brr) - 1 - i); j++ {
			if brr[j]-brr[j+1] > 0 {
				temp := brr[j]
				brr[j] = brr[j+1]
				brr[j+1] = temp
			}
		}
	}
	for i, j := range brr {
		fmt.Println("第", i, "号为的值为:", j)
	}
}

func main() {
	arrays := []int{7,9,6,5,3}
	bubbleSortNo2(arrays)
}