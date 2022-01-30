package main

import "sort"

// Difficulty:
// Medium

// Tags:
// Sorting
// Monotonic Stack

func numberOfWeakCharacters(properties [][]int) int {
	sort.Slice(properties, func(i, j int) bool {
		return properties[i][0] < properties[j][0] ||
			properties[i][0] == properties[j][0] && properties[i][1] > properties[j][1]
	})
	stk := []int{}
	out := 0
	for _, p := range properties {
		for len(stk) > 0 && stk[len(stk)-1] < p[1] {
			stk = stk[:len(stk)-1]
			out++
		}
		stk = append(stk, p[1])
	}
	return out
}
