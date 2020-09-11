package main

/*
  给定一个未排序的整数数组，找出其中没有出现的最小的正整数。
*/

func firstMissingPositive(nums []int) int {
	n := len(nums)
	for i := range nums {
		next := nums[i] - 1
		for 0 <= next && next < n && next != i && nums[next] != nums[i] {
			nums[i], nums[next] = nums[next], nums[i]
			next = nums[i] - 1
		}
	}
	for i := range nums {
		if nums[i] != i+1 {
			return i + 1
		}
	}
	return n + 1
}
