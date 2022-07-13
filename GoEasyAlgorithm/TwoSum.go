package main

import "fmt"

func twoSum(nums []int, target int) (answer []int) {
	// Create a map of array index and array value
	answerMap := map[int]int{}

	// We will change the calculation -> a + b = c change to c - a = b
	for key, value := range nums {
		if v, ok := answerMap[target-value]; ok {
			// Return index about what we need
			answer = []int{v, key}
			break
		}
		// If this value is not what we need, change this value index
		answerMap[value] = key
	}
	return
}

func main() {
	temp := []int{2, 7, 11, 15}
	fmt.Println((twoSum(temp, 9)))
}
