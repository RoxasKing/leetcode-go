package main

func minOperations(nums []int) int {
	out := 0
	n := len(nums)
	for i := 1; i < n; i++ {
		if nums[i] > nums[i-1] {
			continue
		}
		out += nums[i-1] + 1 - nums[i]
		nums[i] = nums[i-1] + 1
	}
	return out
}
