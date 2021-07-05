package main

import "sort"

// Math
func nextPermutation(nums []int) {
	n := len(nums)
	i := n - 1
	for i > 0 && nums[i-1] >= nums[i] {
		i--
	}
	reverse(nums[i:])
	if i == 0 {
		return
	}
	j := sort.Search(n-i, func(j int) bool { return nums[j+i] > nums[i-1] }) + i
	nums[i-1], nums[j] = nums[j], nums[i-1]
}

func reverse(nums []int) {
	n := len(nums)
	for i := 0; i < n>>1; i++ {
		nums[i], nums[n-1-i] = nums[n-1-i], nums[i]
	}
}
