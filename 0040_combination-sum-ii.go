package leetcode

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

func combinationSum2(candidates []int, target int) [][]int {
	sort.Ints(candidates)
	var combination func([]int, int) [][]int
	combination = func(candidates []int, target int) (out [][]int) {
		for i := 0; i < len(candidates); {
			if candidates[i] == target {
				out = append(out, []int{candidates[i]})
			} else if target-candidates[i] >= candidates[i] {
				temp := combination(candidates[i+1:], target-candidates[i])
				if temp != nil {
					for j := range temp {
						out = append(out,
							append([]int{candidates[i]}, temp[j]...))
					}
				}
			}
			for i+1 < len(candidates) && candidates[i+1] == candidates[i] {
				i++
			}
			i++
		}
		return
	}
	return combination(candidates, target)
}
