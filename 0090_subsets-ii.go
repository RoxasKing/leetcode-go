package main

import "sort"

/*
  给定一个可能包含重复元素的整数数组 nums，返回该数组所有可能的子集（幂集）。
  说明：解集不能包含重复的子集。
*/

func subsetsWithDup(nums []int) [][]int {
	sort.Ints(nums)
	var (
		out   [][]int
		recur func([]int, int, int)
	)
	recur = func(tmp []int, size, index int) {
		if len(tmp) == size {
			out = append(out, append(make([]int, 0, size), tmp...))
			return
		}
		for i := index; i < len(nums); i++ {
			tmp = append(tmp, nums[i])
			recur(tmp, size, i+1)
			tmp = tmp[:len(tmp)-1]
			for i+1 < len(nums) && nums[i] == nums[i+1] {
				i++
			}
		}
	}
	for i := 0; i <= len(nums); i++ {
		recur(make([]int, 0, i), i, 0)
	}
	return out
}
