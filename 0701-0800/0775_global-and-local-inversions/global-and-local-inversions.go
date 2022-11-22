package main

// Difficulty:
// Medium

// Tags:
// Math

func isIdealPermutation(nums []int) bool {
	c := 0
	for i := 1; i < len(nums); i++ {
		if nums[i-1] > nums[i] {
			c++
		}
		if i > nums[i] {
			c -= i - nums[i]
		} else if nums[i-1] > nums[i] {
			c--
		}
	}
	return c == 0
}
