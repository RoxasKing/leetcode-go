package main

/*
  给定一个非负整数数组，你最初位于数组的第一个位置。
  数组中的每个元素代表你在该位置可以跳跃的最大长度。
  判断你是否能够到达最后一个位置。
*/

func canJump(nums []int) bool {
	lastPos := len(nums) - 1
	for i := len(nums) - 1; i >= 0; i-- {
		if i+nums[i] >= lastPos {
			lastPos = i
		}
	}
	return lastPos == 0
}

func canJump2(nums []int) bool {
	var lastPos int
	for i := 0; i < len(nums)-1 && i <= lastPos; i++ {
		lastPos = Max(lastPos, nums[i]+i)
	}
	return lastPos >= len(nums)-1
}
