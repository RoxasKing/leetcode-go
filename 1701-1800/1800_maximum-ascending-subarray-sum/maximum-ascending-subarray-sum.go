package main

// Difficulty:
// Easy

func maxAscendingSum(nums []int) int {
	o, sum := nums[0], nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i-1] >= nums[i] {
			sum = 0
		}
		if sum += nums[i]; o < sum {
			o = sum
		}
	}
	return o
}
