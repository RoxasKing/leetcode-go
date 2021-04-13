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

  示例 1：
    输入：candidates = [2,3,6,7], target = 7,
    所求解集为：
    [
      [7],
      [2,2,3]
    ]

  示例 2：
    输入：candidates = [2,3,5], target = 8,
    所求解集为：
    [
      [2,2,2,2],
      [2,3,3],
      [3,5]
    ]

  提示：
    1 <= candidates.length <= 30
    1 <= candidates[i] <= 200
    candidate 中的每个元素都是独一无二的。
    1 <= target <= 500

  来源：力扣（LeetCode）
  链接：https://leetcode-cn.com/problems/combination-sum
  著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

// DFS + Backtracking
func combinationSum(candidates []int, target int) [][]int {
	var out [][]int
	var comb []int
	dfs(candidates, target, &out, &comb, 0)
	return out
}

func dfs(candidates []int, target int, out *[][]int, comb *[]int, index int) {
	if target == 0 {
		tmp := make([]int, len(*comb))
		copy(tmp, *comb)
		*out = append(*out, tmp)
		return
	}
	if index == len(candidates) {
		return
	}
	dfs(candidates, target, out, comb, index+1)
	if target -= candidates[index]; target >= 0 {
		*comb = append(*comb, candidates[index])
		dfs(candidates, target, out, comb, index)
		*comb = (*comb)[:len(*comb)-1]
	}
}

// Dynamic Programming + Hash
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
	base := candidates[0]
	memo := make([][][]int, target-base+1)
	for i := base; i <= target; i++ {
		index := i - base
		for _, candidate := range candidates {
			remain := i - candidate
			if remain < 0 {
				break
			} else if remain == 0 {
				memo[index] = append(memo[index], []int{candidate})
				continue
			}
			if remain -= base; remain < 0 {
				continue
			}
			for _, list := range memo[remain] {
				n := len(list)
				if list[n-1] > candidate {
					continue
				}
				tmp := make([]int, n+1)
				copy(tmp, list)
				tmp[n] = candidate
				memo[index] = append(memo[index], tmp)
			}
		}
	}
	return memo[target-base]
}
