package main

// Difficulty:
// Medium

// Tags:
// Two Pointers

func numSubarrayBoundedMax(nums []int, left int, right int) int {
	o, l, r := 0, -1, -1
	for i, x := range nums {
		if left <= x && x <= right {
			r = i
		} else if x > right {
			l, r = i, -1
		}
		if r != -1 {
			o += r - l
		}
	}
	return o
}
