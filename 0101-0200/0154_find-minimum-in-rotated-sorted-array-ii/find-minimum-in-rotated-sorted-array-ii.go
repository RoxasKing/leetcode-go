package main

// Tags:
// Binary Search
func findMin(nums []int) int {
	l, r := 0, len(nums)-1
	for l < r && nums[l] == nums[r] {
		l++
	}
	return bSearchRotateNum(nums, l, r)
}

func bSearchRotateNum(nums []int, l, r int) int {
	start := l
	for l < r {
		m := (l + r) >> 1
		if nums[m] > nums[m+1] {
			return nums[m+1]
		}
		if nums[l] <= nums[m] {
			l = m + 1
		} else {
			r = m
		}
	}
	return nums[start]
}
