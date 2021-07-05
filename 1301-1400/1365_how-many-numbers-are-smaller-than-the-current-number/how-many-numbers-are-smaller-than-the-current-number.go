package main

import "sort"

func smallerNumbersThanCurrent(nums []int) []int {
	index := make([]int, len(nums))
	for i := range index {
		index[i] = i
	}
	sort.Slice(index, func(i, j int) bool { return nums[index[i]] < nums[index[j]] })
	count := make([]int, len(nums))
	for i := range count {
		if i > 0 && nums[index[i]] == nums[index[i-1]] {
			count[index[i]] = count[index[i-1]]
		} else {
			count[index[i]] = i
		}
	}
	return count
}
