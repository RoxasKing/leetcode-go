package main

// Tags:
// Greedy
func canJump(nums []int) bool {
	n, lastPos := len(nums), 0
	for i := 0; i < n && i <= lastPos && lastPos < n-1; i++ {
		lastPos = Max(lastPos, i+nums[i])
	}
	return lastPos >= n-1
}

func Max(a, b int) int {
	if a < b {
		return b
	}
	return a
}
