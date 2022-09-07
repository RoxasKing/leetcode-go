package main

import "sort"

// Difficulty:
// Medium

// Tags:
// Sorting

func findLongestChain(pairs [][]int) int {
	sort.Slice(pairs, func(i, j int) bool { return pairs[i][1] < pairs[j][1] })
	stk := []int{}
	top := func() int { return len(stk) - 1 }
	for _, p := range pairs {
		if len(stk) > 0 && stk[top()] >= p[0] {
			continue
		}
		stk = append(stk, p[1])
	}
	return len(stk)
}
