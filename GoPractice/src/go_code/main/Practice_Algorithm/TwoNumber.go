package main

import "fmt"

func twoSum(nums []int, target int) []int {
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				return []int{i, j}
			}
		}
	}
	return nil
}

/*
算法转换:1+2=3转变成3-2=1,我们知道预期求3那么利用和减加数=另一个加数的算法来求解
将数组转化成映射
*/
func twoSumMap(nums []int, target int) (answer []int) {
	// Create a map
	tarMap := map[int]int{}

	// Through the array index and value
	for i, num := range nums {
		if v, ok := tarMap[target-num]; ok {
			// Introuding a array and push two number in it
			answer = []int{v, i} // Because this variable introuding in return value
			// Break
			break
		}
		tarMap[num] = i
	}
	return
}

func main() {
	num := []int{5, 3, 6, 2, 1}
	target := 8
	tarNum := twoSum(num, target)
	fmt.Println(tarNum)
}
