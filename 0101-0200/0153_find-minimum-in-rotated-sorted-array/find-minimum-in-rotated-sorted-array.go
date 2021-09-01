package main

// Tags:
// Binary Search

func findMin(nums []int) int {
	for l, r := 0, len(nums)-1; l < r; {
		m := (l + r) >> 1
		if nums[m] > nums[m+1] {
			return nums[m+1]
		}
		if nums[l] < nums[m] {
			l = m + 1
		} else {
			r = m
		}
	}
	return nums[0]
}
