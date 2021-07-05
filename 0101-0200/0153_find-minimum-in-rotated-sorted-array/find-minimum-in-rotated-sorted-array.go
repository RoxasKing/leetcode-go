package main

// Tags:
// Binary Search
func findMin(nums []int) int {
	n := len(nums)
	l, r := 0, n-1
	for l < r {
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
