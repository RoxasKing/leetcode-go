package main

// Difficulty:
// Easy

// Tags:
// Math

func repeatedNTimes(nums []int) int {
	n := len(nums)
	for k := 1; k <= 3; k++ {
		for i := 0; i < n-k; i++ {
			if nums[i] == nums[i+k] {
				return nums[i]
			}
		}
	}
	return -1
}
