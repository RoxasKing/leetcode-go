package main

// Difficulty:
// Hard

// Tags:
// Stack

func longestValidParentheses(s string) int {
	stk := []int{-1}
	top := func() int { return len(stk) - 1 }
	o := 0
	for i, c := range s {
		if c == '(' {
			stk = append(stk, i)
			continue
		}
		if stk = stk[:top()]; len(stk) == 0 {
			stk = append(stk, i)
		}
		o = max(o, i-stk[top()])
	}
	return o
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
