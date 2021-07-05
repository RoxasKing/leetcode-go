package main

import "sort"

// Two Pointers
func minSubArrayLen(target int, nums []int) int {
	out := 0
	n := len(nums)
	for l, r := 0, 0; r < n; r++ {
		target -= nums[r]
		for l < r && target+nums[l] <= 0 {
			target += nums[l]
			l++
		}
		if target <= 0 && (out == 0 || r+1-l < out) {
			out = r + 1 - l
		}
	}
	return out
}

// Binary Seearch + Prefix Sum
func minSubArrayLen2(target int, nums []int) int {
	n := len(nums)
	pSum := make([]int, n+1)
	for i := 0; i < n; i++ {
		pSum[i+1] = pSum[i] + nums[i]
	}

	out := 0
	for i := 0; i < n; i++ {
		j := sort.Search((n+1)-(i+1), func(j int) bool { return pSum[j+(i+1)] >= pSum[i]+target }) + (i + 1)
		if j > n {
			break
		}
		if out == 0 || j-i < out {
			out = j - i
		}
	}
	return out
}
