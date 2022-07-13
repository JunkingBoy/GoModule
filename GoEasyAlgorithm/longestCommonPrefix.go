package main

import "fmt"

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	answer := strs[0]
	length := len(strs)
	for i := 1; i < length; i++ {
		// Always keep the result the shortest string
		answer = lcp(answer, strs[i])
		if len(answer) == 0 {
			break
		}
	}
	return answer
}

func lcp(str1, str2 string) string {
	// The comparison is done in a decreasing way
	length := min(len(str1), len(str2))
	index := 0
	for index < length && str1[index] == str2[index] {
		// The index increasing
		index++
	}
	// Return a slice
	return str1[:index]
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func longestCommonPrefixNo2(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	// The column is locked by the first string
	for i := 0; i < len(strs[0]); i++ {
		// Get each column starting with the second string
		for j := 1; j < len(strs); j++ {
			// Use each column for comparison, If there are unequal characters, Return slice
			if i == len(strs[j]) || strs[j][i] != strs[0][i] {
				// Return slice
				return strs[0][:i]
			}
		}
	}
	return strs[0]
}

func longestCommonPrefixNo3(strs []string) string {
	if len(strs) == 0 {
		return ""
	}

	// Close package
	var lcp func(int, int) string

	// Implement
	lcp = func (start, end int) string {
		if start == end {
			return strs[start]
		}
		mid := (end - start) / 2
		lcpLeft, lcpRight := lcp(start, mid), lcp(mid + 1, end)
		minLength := min(len(lcpLeft), len(lcpRight))
		for i := 0; i < minLength; i++ {
			if lcpLeft[i] != lcpRight[i] {
				return lcpLeft[:i]
			}
		}
		return lcpLeft[: minLength]
	}
	return lcp(0, len(strs) - 1)
}

func main() {
	strs := []string{"flower","flower","flower","flower"}

	fmt.Println(longestCommonPrefixNo3(strs))
}
