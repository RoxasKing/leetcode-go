package main

/*
  给定一组不含重复元素的整数数组 nums，返回该数组所有可能的子集（幂集）。
  说明：解集不能包含重复的子集
*/

// Backtracking
func subsets(nums []int) [][]int {
	var out [][]int
	for i := 0; i <= len(nums); i++ {
		backtrack(nums, i, 0, []int{}, &out)
	}
	return out
}

func backtrack(nums []int, size, start int, set []int, sets *[][]int) {
	if len(set) == size {
		tmp := make([]int, size)
		copy(tmp, set)
		*sets = append(*sets, tmp)
		return
	}
	if start == len(nums) {
		return
	}
	for i := start; i < len(nums); i++ {
		backtrack(nums, size, i+1, append(set, nums[i]), sets)
	}
}
