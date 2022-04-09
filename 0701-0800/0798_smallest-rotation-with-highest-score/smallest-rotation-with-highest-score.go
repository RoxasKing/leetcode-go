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
	o := 0
	for i := 1; i < n; i++ {
		if h[i] += h[i-1] - 1; h[o] > h[i] {
			o = i
		}
	}
	return o
}
