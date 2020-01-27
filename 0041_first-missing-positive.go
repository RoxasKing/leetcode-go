package My_LeetCode_In_Go

import (
	. "My_LeetCode_In_Go/util"
)

/*
  给定一个未排序的整数数组，找出其中没有出现的最小的正整数。
*/

func firstMissingPositive(nums []int) int {
	var containOne bool
	for i := range nums {
		if nums[i] == 1 {
			containOne = true
		} else if nums[i] <= 0 || nums[i] > len(nums) {
			nums[i] = 1
		}
	}
	if !containOne {
		return 1
	}
	for i := range nums {
		num := Abs(nums[i])
		if nums[num-1] > 0 {
			nums[num-1] = -nums[num-1]
		}
	}
	for i := range nums {
		if nums[i] > 0 {
			return i + 1
		}
	}
	return len(nums) + 1
}
