package main

import "fmt"

func main() {
	fmt.Println(isPalindrome(-12321))
}

func isPalindrome(x int) bool {
	// If this varibale < 0 it is 10 times, return false
	if 0 > x || (x%10 == 0 && x != 0) {
		return false
	}

	// Judge the remainder
	revertedNumber := 0
	for x > revertedNumber {
		// If the length of x is odd the last time we will get the number an even length number and an odd length number
		revertedNumber = revertedNumber*10 + x%10
		// x to reduce the ten times
		x /= 10
	}
	// Return
	return x == revertedNumber || x == (revertedNumber/10)
}
