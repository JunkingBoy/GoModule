package main

func removeDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	// Declare a number to record length
	answer := 1

	for fast := 1; fast < len(nums); fast++ {
		if nums[fast] != nums[fast - 1] {
			// Records the value of the current index
			nums[answer] = nums[fast]
			answer++
		}
	}
	return answer
}

func removeDuplicatesNo2(nums []int) int {
		if len(nums) == 0 {
			return 0
		}
		// Double pointer
		left, right := 1, 1

		// Double pointer to move
		for right < len(nums) {
			if nums[right] != nums[right - 1] {
				// Modify nums[left] value
				nums[left] = nums[right]
				// Left pointer move
				left++
			}
			// Right pointer move
			right++
		}
		// Return left pointer
		return left
}