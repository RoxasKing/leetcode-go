package main

// Difficulty:
// Hard

// Tags:
// Prefix Sum

func bestRotation(nums []int) int {
	n := len(nums)
	h := make([]int, n)
	for i, num := range nums {
		h[(i-num+n+1)%n]++
	}
	out := 0
	for i := 1; i < n; i++ {
		h[i] += h[i-1] - 1
		if h[out] > h[i] {
			out = i
		}
	}
	return out
}
