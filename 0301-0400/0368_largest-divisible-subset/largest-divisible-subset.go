package main

import "sort"

// Difficulty:
// Medium

// Tags:
// Dynamic Programming
// Math
// Sorting

func largestDivisibleSubset(nums []int) []int {
	sort.Ints(nums)
	n := len(nums)
	pre := make([]int, n)
	cnt := make([]int, n)
	idx, max := -1, 0
	for i := 0; i < n; i++ {
		pre[i] = -1
		cnt[i] = 1
		for j := i - 1; j >= 0; j-- {
			if nums[i]%nums[j] == 0 && cnt[j]+1 > cnt[i] {
				pre[i] = j
				cnt[i] = cnt[j] + 1
			}
		}
		if cnt[i] > max {
			idx, max = i, cnt[i]
		}
	}
	out := make([]int, max)
	for i := max - 1; i >= 0; i-- {
		out[i] = nums[idx]
		idx = pre[idx]
	}
	return out
}
