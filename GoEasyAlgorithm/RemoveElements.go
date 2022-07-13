package main

import "container/ring"

func removeElement(nums []int, val int) int {
	left, right := 0, len(nums)

	for left < right {
		if nums[left] == val {
			nums[left] = nums[right - 1]
			// Right move to left
			right--
		}else {
			left++			
		}
	}
	return left
}