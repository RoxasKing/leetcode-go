package main

// Difficulty:
// Medium

// Tags:
// Prefix Sum
// Sliding Window

func maximumUniqueSubarray(nums []int) int {
	n := len(nums)
	pSum := make([]int, n+1)
	for i, x := range nums {
		pSum[i+1] = pSum[i] + x
	}
	h := [1e4 + 1]bool{}
	o := 0
	for l, r := 0, 0; r < n; r++ {
		for ; h[nums[r]]; l++ {
			h[nums[l]] = false
		}
		h[nums[r]] = true
		if o < pSum[r+1]-pSum[l] {
			o = pSum[r+1] - pSum[l]
		}
	}
	return o
}
