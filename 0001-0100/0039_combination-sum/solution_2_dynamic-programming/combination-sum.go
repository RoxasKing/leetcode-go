package main

import "sort"

// Difficulty:
// Medium

// Tags:
// Dynamic Programming

func combinationSum(candidates []int, target int) [][]int {
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
