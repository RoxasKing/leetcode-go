package main

import (
	"sort"
)

// Tags:
// Backtracking
func combinationSum2(candidates []int, target int) [][]int {
	sort.Ints(candidates)
	out := [][]int{}
	dfs(candidates, target, len(candidates), 0, []int{}, &out)
	return out
}

func dfs(candidates []int, target int, n, i int, cur []int, out *[][]int) {
	if target == 0 {
		tmp := make([]int, len(cur))
		copy(tmp, cur)
		*out = append(*out, tmp)
		return
	}

	if i == n {
		return
	}

	if target-candidates[i] >= 0 {
		target -= candidates[i]
		cur = append(cur, candidates[i])
		dfs(candidates, target, n, i+1, cur, out)
		cur = cur[:len(cur)-1]
		target += candidates[i]
	}

	for i+1 < n && candidates[i+1] == candidates[i] {
		i++
	}
	dfs(candidates, target, n, i+1, cur, out)
}
