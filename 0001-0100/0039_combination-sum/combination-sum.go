package main

import (
	"sort"
)

/*
  Given an array of distinct integers candidates and a target integer target, return a list of all unique combinations of candidates where the chosen numbers sum to target. You may return the combinations in any order.

  The same number may be chosen from candidates an unlimited number of times. Two combinations are unique if the frequency of at least one of the chosen numbers is different.

  It is guaranteed that the number of unique combinations that sum up to target is less than 150 combinations for the given input.

  Example 1:
    Input: candidates = [2,3,6,7], target = 7
    Output: [[2,2,3],[7]]
    Explanation:
      2 and 3 are candidates, and 2 + 2 + 3 = 7. Note that 2 can be used multiple times.
      7 is a candidate, and 7 = 7.
      These are the only two combinations.

  Example 2:
    Input: candidates = [2,3,5], target = 8
    Output: [[2,2,2,2],[2,3,3],[3,5]]

  Example 3:
    Input: candidates = [2], target = 1
    Output: []

  Example 4:
    Input: candidates = [1], target = 1
    Output: [[1]]

  Example 5:
    Input: candidates = [1], target = 2
    Output: [[1,1]]

  Constraints:
    1. 1 <= candidates.length <= 30
    2. 1 <= candidates[i] <= 200
    3. All elements of candidates are distinct.
    4. 1 <= target <= 500

  来源：力扣（LeetCode）
  链接：https://leetcode-cn.com/problems/combination-sum
  著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

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
