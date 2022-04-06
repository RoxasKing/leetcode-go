package main

// Difficulty:
// Easy

// Tags:
// Monotonic Stack

func nextGreatestLetter(letters []byte, target byte) byte {
	stk := []int{}
	top := func() int { return len(stk) - 1 }
	for _, letter := range letters {
		x := int(letter - 'a')
		if letter <= target {
			x += 26
		}
		for ; len(stk) > 0 && stk[top()] >= x; stk = stk[:top()] {
		}
		stk = append(stk, x)
	}
	return byte(stk[0]%26) + 'a'
}
