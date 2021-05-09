package main

import "sort"

/*
  Given an integer array nums that may contain duplicates, return all possible subsets (the power set).

  The solution set must not contain duplicate subsets. Return the solution in any order.

  Example 1:
    Input: nums = [1,2,2]
    Output: [[],[1],[1,2],[1,2,2],[2],[2,2]]

  Example 2:
    Input: nums = [0]
    Output: [[],[0]]

  Constraints:
    1 <= nums.length <= 10
    -10 <= nums[i] <= 10

  来源：力扣（LeetCode）
  链接：https://leetcode-cn.com/problems/subsets-ii
  著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

// Backtracking
func subsetsWithDup(nums []int) [][]int {
	sort.Ints(nums)
	out := [][]int{}
	for size := 0; size <= len(nums); size++ {
		cur := make([]int, 0, size)
		dfs(nums, 0, size, cur, &out)
	}
	return out
}

func dfs(nums []int, idx, size int, subset []int, out *[][]int) {
	if len(subset) == size {
		tmp := make([]int, len(subset))
		copy(tmp, subset)
		*out = append(*out, tmp)
		return
	}

	if idx == len(nums) {
		return
	}

	subset = append(subset, nums[idx])
	dfs(nums, idx+1, size, subset, out)
	subset = subset[:len(subset)-1]

	next := idx + 1
	for next < len(nums) && nums[next] == nums[idx] {
		next++
	}
	dfs(nums, next, size, subset, out)
}
