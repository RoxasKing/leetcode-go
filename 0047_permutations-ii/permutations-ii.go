package main

import (
	"sort"
)

/*
  给定一个可包含重复数字的序列，返回所有不重复的全排列。
*/

// Backtracking
func permuteUnique(nums []int) [][]int {
	sort.Ints(nums)
	n := len(nums)
	used := make([]bool, n)
	cur := make([]int, 0, n)
	var out [][]int
	var backtrack func(int)
	backtrack = func(index int) {
		if index == n {
			tmp := make([]int, n)
			copy(tmp, cur)
			out = append(out, tmp)
			return
		}
		for i := 0; i < n; i++ {
			if used[i] || i > 0 && !used[i-1] && nums[i] == nums[i-1] { // 保证重复数字按顺序取
				continue
			}
			cur = append(cur, nums[i])
			used[i] = true
			backtrack(index + 1)
			used[i] = false
			cur = cur[:len(cur)-1]
		}
	}
	backtrack(0)
	return out
}
