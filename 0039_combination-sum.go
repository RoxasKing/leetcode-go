package leetcode

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

func combinationSum(candidates []int, target int) [][]int {
	sort.Ints(candidates)
	var combination func([]int, int) [][]int
	combination = func(candidates []int, target int) (out [][]int) {
		for i := range candidates {
			if candidates[i] == target {
				out = append(out, []int{candidates[i]})
			} else if target-candidates[i] >= candidates[i] {
				temp := combination(candidates[i:], target-candidates[i])
				if temp == nil {
					continue
				}
				for j := range temp {
					out = append(out, append([]int{candidates[i]}, temp[j]...))
				}
			}
		}
		return
	}
	return combination(candidates, target)
}
