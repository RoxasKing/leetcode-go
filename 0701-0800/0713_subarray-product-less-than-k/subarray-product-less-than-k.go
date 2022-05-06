package main

// Difficulty:
// Medium

// Tags:
// Sliding Window

func numSubarrayProductLessThanK(nums []int, k int) int {
	out := 0
	for l, r, x := 0, 0, 1; r < len(nums); r++ {
		x *= nums[r]
		for ; l <= r && x >= k; l++ {
			x /= nums[l]
		}
		out += r + 1 - l
	}
	return out
}
