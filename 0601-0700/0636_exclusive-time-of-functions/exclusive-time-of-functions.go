package main

import (
	"strconv"
	"strings"
)

// Difficulty:
// Medium

// Tags:
// Stack

func exclusiveTime(n int, logs []string) []int {
	stk := [][3]int{}
	top := func() int { return len(stk) - 1 }
	o := make([]int, n)
	for _, log := range logs {
		a := strings.Split(log, ":")
		i, _ := strconv.Atoi(a[0])
		t, _ := strconv.Atoi(a[2])
		if a[1][0] == 's' {
			stk = append(stk, [3]int{i, t, 0})
			continue
		}
		p := stk[top()]
		stk = stk[:top()]
		o[i] += t + 1 - p[1] - p[2]
		if len(stk) > 0 {
			stk[top()][2] += t + 1 - p[1]
		}
	}
	return o
}
