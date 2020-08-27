package main

/*
  给定一个没有重复数字的序列，返回其所有可能的全排列。
*/

// Backtracking
func permute(nums []int) [][]int {
	var out [][]int
	var backtrack func(int)
	backtrack = func(index int) {
		if index == len(nums) {
			tmp := make([]int, len(nums))
			copy(tmp, nums)
			out = append(out, tmp)
			return
		}
		for i := index; i < len(nums); i++ {
			nums[i], nums[index] = nums[index], nums[i]
			backtrack(index + 1)
			nums[i], nums[index] = nums[index], nums[i]
		}
	}
	backtrack(0)
	return out
}
