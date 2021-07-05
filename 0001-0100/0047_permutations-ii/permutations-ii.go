package main

import (
	"sort"
)

// Tags:
// Backtracking
func permuteUnique(nums []int) [][]int {
	var out [][]int
	n := len(nums)
	sort.Ints(nums)
	used := make([]bool, n)
	cur := make([]int, 0, n)
	backtrack(nums, n, &out, used, &cur, 0)
	return out
}

func backtrack(nums []int, n int, out *[][]int, used []bool, cur *[]int, index int) {
	if index == n {
		tmp := make([]int, n)
		copy(tmp, *cur)
		*out = append(*out, tmp)
		return
	}
	for i := 0; i < n; i++ {
		if used[i] || i > 0 && nums[i] == nums[i-1] && !used[i-1] {
			continue
		}
		*cur = append(*cur, nums[i])
		used[i] = true
		backtrack(nums, n, out, used, cur, index+1)
		used[i] = false
		*cur = (*cur)[:len(*cur)-1]
	}
}
