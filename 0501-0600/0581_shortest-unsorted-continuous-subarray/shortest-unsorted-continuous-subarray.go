package main

// Difficulty:
// Medium

// Tags:
// Two Pointers

func findUnsortedSubarray(nums []int) int {
	l, r := -1, -1
	n := len(nums)
	for min, max, i := 1<<63-1, -1<<63, 0; i < n; i++ {
		if max > nums[i] {
			r = i
		} else {
			max = nums[i]
		}
		if min < nums[n-1-i] {
			l = n - 1 - i
		} else {
			min = nums[n-1-i]
		}
	}
	if r == -1 {
		return 0
	}
	return r - l + 1
}
