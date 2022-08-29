package main

// Difficulty:
// Easy

func shuffle(nums []int, n int) []int {
	o := make([]int, n*2)
	for i := 0; i < n; i++ {
		o[i*2] = nums[i]
		o[i*2+1] = nums[n+i]
	}
	return o
}
