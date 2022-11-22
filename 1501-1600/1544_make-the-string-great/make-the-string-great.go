package main

// Difficulty:
// Easy

// Tags:
// Stack

func makeGood(s string) string {
	stk := []rune{}
	top := func() int { return len(stk) - 1 }
	for _, c := range s {
		if len(stk) > 0 && abs(stk[top()]-c) == 32 {
			stk = stk[:top()]
			continue
		}
		stk = append(stk, c)
	}
	return string(stk)
}

func abs(x rune) rune {
	if x < 0 {
		return -x
	}
	return x
}
