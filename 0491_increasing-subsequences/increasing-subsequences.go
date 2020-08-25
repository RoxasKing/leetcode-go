package main

/*
  给定一个整型数组, 你的任务是找到所有该数组的递增子序列，递增子序列的长度至少是2。

  说明:
    给定数组的长度不会超过15。
    数组中的整数范围是 [-100,100]。
    给定数组中可能包含重复数字，相等的数字应该被视为递增的一种情况。
*/

// Backtracking
func findSubsequences(nums []int) [][]int {
	var out [][]int
	var cur []int
	backtrack(nums, &out, &cur, 0)
	return out
}

func backtrack(nums []int, out *[][]int, cur *[]int, index int) {
	if index == len(nums) {
		if len(*cur) >= 2 {
			tmp := make([]int, len(*cur))
			copy(tmp, *cur)
			*out = append(*out, tmp)
		}
		return
	}
	if len(*cur) == 0 || nums[index] >= (*cur)[len(*cur)-1] {
		*cur = append(*cur, nums[index])
		backtrack(nums, out, cur, index+1)
		*cur = (*cur)[:len(*cur)-1]
	}
	if len(*cur) == 0 || nums[index] != (*cur)[len(*cur)-1] {
		backtrack(nums, out, cur, index+1)
	}
}
