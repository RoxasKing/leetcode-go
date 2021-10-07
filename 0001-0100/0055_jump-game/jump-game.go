package main

// Difficulty:
// Medium

// Tags:
// Greedy

func canJump(nums []int) bool {
	n, r := len(nums), 0
	for i := 0; i <= r && r < n-1; i++ {
		if i+nums[i] > r {
			r = i + nums[i]
		}
	}
	return r >= n-1
}
