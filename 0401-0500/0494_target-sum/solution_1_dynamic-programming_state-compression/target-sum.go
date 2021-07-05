package main

import "math/bits"

// Dynamic Programming
// State Compression

func findTargetSumWays(nums []int, target int) int {
	n := len(nums)
	f := make([]int, 1<<n)
	for _, num := range nums {
		f[0] += num
	}
	if f[0] < target {
		return 0
	}
	out := 0
	if f[0] == target {
		out++
	}
	for i := 1; i < 1<<n; i++ {
		j := bits.TrailingZeros(uint(i))
		f[i] = f[i-(1<<j)] - 2*nums[j]
		if f[i] == target {
			out++
		}
	}
	return out
}
