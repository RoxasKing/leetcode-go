package main

// Tags:
// Stack
func longestValidParentheses(s string) int {
	out := 0
	stk := []int{-1}
	for i := range s {
		if s[i] == '(' {
			stk = append(stk, i)
			continue
		}

		if len(stk) == 1 {
			stk[0] = i
			continue
		}

		stk = stk[:len(stk)-1]
		out = Max(out, i-stk[len(stk)-1])
	}

	return out
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
