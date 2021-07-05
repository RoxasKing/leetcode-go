package main

import "sort"

func findUnsortedSubarray(nums []int) int {
	n := len(nums)
	sorted := make([]int, n)
	copy(sorted, nums)
	sort.Ints(sorted)
	l, r := 0, n-1
	for l < n && sorted[l] == nums[l] {
		l++
	}
	for r >= l && sorted[r] == nums[r] {
		r--
	}
	return r + 1 - l
}

// Two Pointers
func findUnsortedSubarray2(nums []int) int {
	n := len(nums)
	min, max := 1<<31-1, -1<<31
	for i := 0; i < n-1; i++ {
		if nums[i] > nums[i+1] {
			max = Max(max, nums[i])
			min = Min(min, nums[i+1])
		}
	}

	l, r := 0, 0
	for i := 0; i < n; i++ {
		if nums[i] > min {
			l = i
			break
		}
	}
	for i := n - 1; i >= 0; i-- {
		if nums[i] < max {
			r = i
			break
		}
	}

	if l >= r {
		return 0
	}
	return r + 1 - l
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
