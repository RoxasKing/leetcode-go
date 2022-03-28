package main

import "strconv"

// Difficulty:
// Easy

// Tags:
// Stack

func calPoints(ops []string) int {
	stk := []int{}
	top := func() int { return len(stk) - 1 }
	for _, op := range ops {
		switch op[0] {
		case '+':
			stk = append(stk, stk[top()]+stk[top()-1])
		case 'D':
			stk = append(stk, stk[top()]<<1)
		case 'C':
			stk = stk[:top()]
		default:
			x, _ := strconv.Atoi(op)
			stk = append(stk, x)
		}
	}
	sum := 0
	for _, x := range stk {
		sum += x
	}
	return sum
}
