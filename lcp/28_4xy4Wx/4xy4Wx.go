package main

import "sort"

// Two Pointers
func purchasePlans(nums []int, target int) int {
	sort.Ints(nums)
	l, r := 0, len(nums)-1
	out := 0
	for l < r {
		sum := nums[l] + nums[r]
		if sum <= target {
			out += r - l
			out %= 1e9 + 7
			l++
		} else {
			r--
		}
	}
	return out
}

// Binary Search
func purchasePlans2(nums []int, target int) int {
	sort.Ints(nums)
	out := 0
	for i := range nums {
		if nums[i] >= target {
			break
		}
		j := sort.Search(len(nums), func(j int) bool { return nums[j] > target-nums[i] })
		if i >= j-1 {
			break
		}
		out += j - 1 - i
		out %= 1e9 + 7
	}
	return out
}
