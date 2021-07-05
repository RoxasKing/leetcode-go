package main

func buildArray(nums []int) []int {
	n := len(nums)
	out := make([]int, n)
	for i := range nums {
		out[i] = nums[nums[i]]
	}
	return out
}
