package main

import "strings"

// Difficulty:
// Easy

// Tags:
// Stack

func removeOuterParentheses(s string) string {
	arr := []string{}
	stk := []int{-1}
	top := func() int { return len(stk) - 1 }
	cnt := 0
	for i, c := range s {
		if c == '(' {
			stk = append(stk, i)
			cnt++
			continue
		}
		if cnt--; cnt > 0 {
			stk = stk[:top()]
			continue
		}
		l, r := stk[top()]+1, i
		arr = append(arr, s[l:r])
		stk = stk[:top()]
	}
	return strings.Join(arr, "")
}
