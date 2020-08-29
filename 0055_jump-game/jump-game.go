package main

/*
  给定一个非负整数数组，你最初位于数组的第一个位置。
  数组中的每个元素代表你在该位置可以跳跃的最大长度。
  判断你是否能够到达最后一个位置。
*/

// Greedy algorithm
func canJump(nums []int) bool {
	n := len(nums)
	var lastPos int
	for i := 0; i < n-1 && i <= lastPos && lastPos < n-1; i++ {
		lastPos = Max(lastPos, nums[i]+i)
	}
	return lastPos >= n-1
}

func Max(a, b int) int {
	if a < b {
		return b
	}
	return a
}
