package main

import "sort"

// Difficulty:
// Medium

// Tags:
// Backtracking
// Bitmask

func canPartitionKSubsets(nums []int, k int) bool {
	sort.Ints(nums)
	n := len(nums)
	sum := 0
	for _, num := range nums {
		sum += num
	}
	if sum%k != 0 || nums[n-1] > sum/k {
		return false
	}
	avg := sum / k
	b := make([]int, k)
	for i := range b {
		b[i] = avg
	}
	var backtrack func(idx int) bool
	backtrack = func(idx int) bool {
		if idx < 0 {
			return true
		}
		for i := range b {
			if b[i] == nums[idx] || b[i]-nums[idx] >= nums[0] {
				b[i] -= nums[idx]
				if backtrack(idx - 1) {
					return true
				}
				b[i] += nums[idx]
			}
			if b[i] == avg {
				break
			}
		}
		return false
	}
	return backtrack(n - 1)
}
