package main

// Difficulty:
// Medium

// Tags:
// Prefix Sum

func productExceptSelf(nums []int) []int {
	n := len(nums)
	out := make([]int, n)
	out[n-1] = 1
	for i := n - 2; i >= 0; i-- {
		out[i] = nums[i+1] * out[i+1]
	}
	for i, mul := 0, 1; i < n; i++ {
		out[i] *= mul
		mul *= nums[i]
	}
	return out
}
