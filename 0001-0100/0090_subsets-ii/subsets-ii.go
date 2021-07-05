package main

import "sort"

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
