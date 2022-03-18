package main

// Difficulty:
// Medium

// Tags:
// Stack

func scoreOfParentheses(s string) int {
	stk := []int{0}
	top := func() int { return len(stk) - 1 }
	for i := range s {
		if c := s[i]; c == '(' {
			stk = append(stk, 0)
		} else {
			k := stk[top()]
			stk = stk[:top()]
			stk[top()] += Max(k<<1, 1)
		}
	}
	return stk[0]
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
