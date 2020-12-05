package main

import "sort"

/*
  给定一个可能包含重复元素的整数数组 nums，返回该数组所有可能的子集（幂集）。
  说明：解集不能包含重复的子集。
*/

// DFS
func subsetsWithDup(nums []int) [][]int {
	sort.Ints(nums)
	out := [][]int{}
	for i := 0; i <= len(nums); i++ {
		cur := make([]int, 0, i)
		dfs(nums, i, 0, &cur, &out)
	}
	return out
}

func dfs(nums []int, size, start int, cur *[]int, out *[][]int) {
	if len(*cur) == size {
		tmp := make([]int, len(*cur))
		copy(tmp, *cur)
		*out = append(*out, tmp)
		return
	}
	if start == len(nums) {
		return
	}

	*cur = append(*cur, nums[start])
	dfs(nums, size, start+1, cur, out)
	*cur = (*cur)[:len(*cur)-1]

	next := start
	for next < len(nums) && nums[next] == nums[start] {
		next++
	}
	dfs(nums, size, next, cur, out)
}
