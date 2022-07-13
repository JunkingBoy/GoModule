package main

func isValid(s string) bool {
	// If s.length % 2 == 1 -> false
	if len(s) % 2 == 1 {
		return false
	}

	// Create a map and use stack
	pairs := map[byte]byte{
		')': '(',
		']': '[',
		'}': '{',
	}

	// Create a stack
	stacks := []byte{}

	for i := 0; i < len(s); i++ {
		if pairs[s[i]] > 0 {
			if len(stacks) == 0 || stacks[len(stacks) - 1] != pairs[s[i]] {
				return false
			}
			stacks = stacks[:len(stacks) - 1]
		}else {
			// Input into stack
			stacks = append(stacks, s[i])
		}
	}
	return len(stacks) == 0
}