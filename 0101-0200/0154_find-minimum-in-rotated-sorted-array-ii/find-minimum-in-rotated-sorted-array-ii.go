package main

// Difficulty:
// Hard

// Tags:
// Binary Search

func findMin(nums []int) int {
	l, r := 0, len(nums)-1
	for ; l < r && nums[l] == nums[r]; r-- {
	}
	if nums[l] <= nums[r] {
		return nums[l]
	}
	for l < r {
		m := l + (r-l)>>1
		if nums[m] > nums[m+1] {
			return nums[m+1]
		}
		if nums[l] <= nums[m] {
			l = m + 1
		} else {
			r = m
		}
	}
	return nums[l]
}
