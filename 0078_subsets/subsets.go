package main

/*
  给定一组不含重复元素的整数数组 nums，返回该数组所有可能的子集（幂集）。
  说明：解集不能包含重复的子集
*/

// Backtracking
func subsets(nums []int) [][]int {
	out := [][]int{{}}
	var cur []int
	var backtrack func(int, int)
	backtrack = func(limit, start int) {
		if len(cur) == limit {
			tmp := make([]int, len(cur))
			copy(tmp, cur)
			out = append(out, tmp)
			return
		}
		for i := start; i < len(nums); i++ {
			cur = append(cur, nums[i])
			backtrack(limit, i+1)
			cur = cur[:len(cur)-1]
		}
	}
	for i := 1; i <= len(nums); i++ {
		backtrack(i, 0)
	}
	return out
}
