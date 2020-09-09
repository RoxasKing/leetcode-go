package main

import (
	"sort"
)

/*
  给定一个无重复元素的数组 candidates 和一个目标数 target ，找出 candidates 中所有可以使数字和为 target 的组合。
  candidates 中的数字可以无限制重复被选取。
  说明：
    所有数字（包括 target）都是正整数。
    解集不能包含重复的组合。
*/

// DFS + Backtracking
func combinationSum(candidates []int, target int) [][]int {
	var out [][]int
	var comb []int
	var dfs func(int, int)
	dfs = func(index, target int) {
		if index == len(candidates) {
			return
		}
		if target == 0 {
			tmp := make([]int, len(comb))
			copy(tmp, comb)
			out = append(out, tmp)
			return
		}
		dfs(index+1, target)
		if target-candidates[index] >= 0 {
			comb = append(comb, candidates[index])
			dfs(index, target-candidates[index])
			comb = comb[:len(comb)-1]
		}
	}
	dfs(0, target)
	return out
}

// Hash
func combinationSum2(candidates []int, target int) [][]int {
	sort.Ints(candidates)
	last := len(candidates) - 1
	for last >= 0 && candidates[last] > target {
		last--
	}
	candidates = candidates[:last+1]
	if len(candidates) == 0 {
		return nil
	}
	memo := make(map[int][][]int, target+1)
	for i := candidates[0]; i <= target; i++ {
		for _, candidate := range candidates {
			if i-candidate < 0 {
				break
			}
			if lists, ok := memo[i-candidate]; ok {
				for _, list := range lists {
					if list[len(list)-1] > candidate {
						continue
					}
					tmp := make([]int, len(list)+1)
					copy(tmp, list)
					tmp[len(list)] = candidate
					memo[i] = append(memo[i], tmp)
				}
			}
			if i == candidate {
				memo[i] = append(memo[i], []int{candidate})
			}
		}
	}
	return memo[target]
}
