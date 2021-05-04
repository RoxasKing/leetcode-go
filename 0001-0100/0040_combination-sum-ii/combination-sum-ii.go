package main

import (
	"sort"
)

/*
  Given a collection of candidate numbers (candidates) and a target number (target), find all unique combinations in candidates where the candidate numbers sum to target.

  Each number in candidates may only be used once in the combination.

  Note: The solution set must not contain duplicate combinations.

  Example 1:
    Input: candidates = [10,1,2,7,6,1,5], target = 8
    Output:
      [
      [1,1,6],
      [1,2,5],
      [1,7],
      [2,6]
      ]

  Example 2:
    Input: candidates = [2,5,2,1,2], target = 5
    Output:
      [
      [1,2,2],
      [5]
      ]

  Constraints:
    1. 1 <= candidates.length <= 100
    2. 1 <= candidates[i] <= 50
    3. 1 <= target <= 30

  来源：力扣（LeetCode）
  链接：https://leetcode-cn.com/problems/combination-sum-ii
  著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

// DFS + Backtracking
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
