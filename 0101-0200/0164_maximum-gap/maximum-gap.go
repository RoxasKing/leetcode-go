package main

import (
	"sort"
)

func maximumGap(nums []int) int {
	if len(nums) < 2 {
		return 0
	}
	sort.Ints(nums)
	var max int
	for i := 1; i < len(nums); i++ {
		max = Max(max, nums[i]-nums[i-1])
	}
	return max
}

// Radix Sort
func maximumGap2(nums []int) int {
	if len(nums) < 2 {
		return 0
	}
	maxNum := -1 << 31
	for _, num := range nums {
		maxNum = Max(maxNum, num)
	}
	n, buckets := len(nums), 10
	tmps := make([]int, n)
	for exp := 1; maxNum/exp > 0; exp *= 10 {
		count := make([]int, buckets)
		for _, num := range nums {
			count[(num/exp)%10]++
		}
		for i := 1; i < buckets; i++ {
			count[i] += count[i-1]
		}
		for i := n - 1; i >= 0; i-- {
			index := (nums[i] / exp) % 10
			count[index]--
			tmps[count[index]] = nums[i]
		}
		copy(nums, tmps)
	}

	out := 0
	for i := 1; i < n; i++ {
		out = Max(out, nums[i]-nums[i-1])
	}
	return out
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
