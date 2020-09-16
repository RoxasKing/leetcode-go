package main

import (
	"sort"
)

/*
  给定一个数组 candidates 和一个目标数 target ，找出 candidates 中所有可以使数字和为 target 的组合。
  candidates 中的每个数字在每个组合中只能使用一次。
  说明：
    所有数字（包括目标数）都是正整数。
    解集不能包含重复的组合。
*/

// Backtracking
func combinationSum2(candidates []int, target int) [][]int {
	if len(candidates) == 0 {
		return nil
	}
	sort.Ints(candidates)
	var out [][]int
	var combination []int
	var backtrack func(int, int)
	backtrack = func(index, target int) {
		if target == 0 {
			tmp := make([]int, len(combination))
			copy(tmp, combination)
			out = append(out, tmp)
			return
		}
		if index == len(candidates) {
			return
		}

		next := index
		for next < len(candidates) && candidates[next] == candidates[index] {
			next++
		}

		backtrack(next, target)

		curSize := len(combination)
		for i := 0; i < next-index && candidates[index] <= target; i++ {
			combination = append(combination, candidates[index])
			target -= candidates[index]
			backtrack(next, target)
		}
		combination = combination[:curSize]
	}
	backtrack(0, target)
	return out
}
