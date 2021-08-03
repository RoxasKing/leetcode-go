package main

// Tags:
// Two Pointers

func findUnsortedSubarray(nums []int) int {
	n := len(nums)
	min, max := 1<<63-1, -1<<63
	l, r := -1, -1
	for i, num := range nums {
		if max > num {
			r = i
		} else {
			max = num
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
