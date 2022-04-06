package main

import "sort"

// Difficulty:
// Medium

// Tags:
// Math

func nextPermutation(nums []int) {
	n := len(nums)
	i := n - 1
	for ; i > 0 && nums[i-1] >= nums[i]; i-- {
	}
	for l, r := i, n-1; l < r; l, r = l+1, r-1 {
		nums[l], nums[r] = nums[r], nums[l]
	}
	if i == 0 {
		return
	}
	j := sort.Search(n-i, func(j int) bool { return nums[j+i] > nums[i-1] }) + i
	nums[i-1], nums[j] = nums[j], nums[i-1]
}
