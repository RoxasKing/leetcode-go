package main

// Difficulty:
// Easy

// Tags:
// Stack

func removeDuplicates(s string) string {
	stk := []rune{}
	top := func() int { return len(stk) - 1 }
	for _, c := range s {
		if len(stk) > 0 && stk[top()] == c {
			stk = stk[:top()]
			continue
		}
		stk = append(stk, c)
	}
	return string(stk)
}
