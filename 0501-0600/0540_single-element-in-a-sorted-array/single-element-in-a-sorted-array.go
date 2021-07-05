package main

// Tags:
// Binary Search
func singleNonDuplicate(nums []int) int {
	n := len(nums)
	l, r := 0, n-1
	for l < r {
		m := (l + r) >> 1
		if m&1 == 0 && nums[m] == nums[m+1] || m&1 == 1 && nums[m-1] == nums[m] {
			l = m + 1
		} else {
			r = m
		}
	}
	return nums[l]
}
