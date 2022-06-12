package main

// Difficulty:
// Medium

// Tags:
// Sliding Window

func minOperations(nums []int, x int) int {
	sum := 0
	for _, x := range nums {
		sum += x
	}
	if sum < x {
		return -1
	}
	x = sum - x
	n := len(nums)
	o := -1
	for l, r := 0, 0; r < n; r++ {
		x -= nums[r]
		for ; x < 0; l++ {
			x += nums[l]
		}
		if x == 0 && o < r+1-l {
			o = r + 1 - l
		}
	}
	if o == -1 {
		return -1
	}
	return n - o
}
