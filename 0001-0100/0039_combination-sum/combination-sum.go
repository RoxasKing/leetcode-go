package main

import (
	"sort"
)

// Tags:
// Backtracking
func combinationSum(candidates []int, target int) [][]int {
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
		cur = append(cur, candidates[i])
		target -= candidates[i]
		dfs(candidates, target, n, i, cur, out)
		target += candidates[i]
		cur = cur[:len(cur)-1]
	}

	dfs(candidates, target, n, i+1, cur, out)
}

// Dynamic Programming
func combinationSum2(candidates []int, target int) [][]int {
	dp := make([][][]int, target+1)
	dp[0] = [][]int{{}}
	sort.Ints(candidates)

	for i := 1; i <= target; i++ {
		for _, candidate := range candidates {
			if candidate > i {
				break
			}
			for _, arr := range dp[i-candidate] {
				if len(arr) > 0 && arr[len(arr)-1] > candidate {
					break
				}
				tmp := make([]int, len(arr)+1)
				copy(tmp[:len(arr)], arr)
				tmp[len(arr)] = candidate
				dp[i] = append(dp[i], tmp)
			}
		}
	}

	return dp[target]
}
