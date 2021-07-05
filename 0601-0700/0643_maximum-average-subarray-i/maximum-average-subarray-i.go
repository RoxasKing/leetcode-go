package main

import "math"

func findMaxAverage(nums []int, k int) float64 {
	n := len(nums)
	out := -10001.0
	sum := 0
	for i := 0; i < n; i++ {
		sum += nums[i]
		if i > k-1 {
			sum -= nums[i-k]
		}
		if i >= k-1 {
			out = math.Max(out, float64(sum)/float64(k))
		}
	}
	return out
}
