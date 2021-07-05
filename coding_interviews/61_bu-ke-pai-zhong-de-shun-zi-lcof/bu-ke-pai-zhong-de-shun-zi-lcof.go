package main

import "sort"

func isStraight(nums []int) bool {
	sort.Ints(nums)
	cnt := 0
	for len(nums) > 0 && nums[0] == 0 {
		cnt++
		nums = nums[1:]
	}
	for i := 1; i < len(nums); i++ {
		if nums[i] == nums[i-1] {
			return false
		}
		cnt -= nums[i] - nums[i-1] - 1
	}
	return cnt >= 0
}
