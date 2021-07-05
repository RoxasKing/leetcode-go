package main

// Tags:
// Binary Search
func missingNumber(nums []int) int {
	l, r := 0, len(nums)-1
	for l <= r {
		m := l + (r-l)>>1
		if nums[m] == m {
			l = m + 1
		} else {
			if m == 0 || nums[m-1] == m-1 {
				return m
			}
			r = m - 1
		}
	}
	return l
}
