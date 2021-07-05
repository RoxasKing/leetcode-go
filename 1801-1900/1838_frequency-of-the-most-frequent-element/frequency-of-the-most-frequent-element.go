package main

import "sort"

// Sliding Window
func maxFrequency(nums []int, k int) int {
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] > nums[j]
	})

	n := len(nums)
	out := 0

	for l, r := 0, 0; r < n; r++ {
		k -= nums[l] - nums[r]
		for l < r && k < 0 {
			k += (r - l) * (nums[l] - nums[l+1])
			l++
		}
		out = Max(out, r+1-l)
	}

	return out
}

// Binary Search
func maxFrequency2(nums []int, k int) int {
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] > nums[j]
	})
	n := len(nums)
	pSum := make([]int, n+1)
	for i := 0; i < n; i++ {
		pSum[i+1] = pSum[i] + nums[i]
	}
	out := 0
	for i := 1; i <= n; i++ {
		j := sort.Search(n+1-(i+1), func(j int) bool {
			j += i + 1
			return (pSum[j]-pSum[i-1]+k)/(j-i+1) < nums[i-1]
		}) + (i + 1)
		out = Max(out, j-i)
	}
	return out
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
