package main

// Tags:
// Binary Search
func searchRange(nums []int, target int) []int {
	n := len(nums)
	l := bSearch(nums, target, 0, n-1, true)
	if l == -1 {
		return []int{-1, -1}
	}
	r := bSearch(nums, target, l, n-1, false)
	return []int{l, r}
}

func bSearch(nums []int, target, l, r int, leftBond bool) int {
	for l <= r {
		m := l + (r-l)>>1
		if nums[m] < target {
			l = m + 1
		} else if nums[m] > target {
			r = m - 1
		} else if leftBond {
			if m == 0 || nums[m-1] < target {
				return m
			}
			r = m - 1
		} else {
			if m == r || nums[m+1] > target {
				return m
			}
			l = m + 1
		}
	}
	return -1
}
