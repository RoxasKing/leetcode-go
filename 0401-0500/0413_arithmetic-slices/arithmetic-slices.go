package main

// Difficulty:
// Medium

// Tags:
// Dynamic Programming

func numberOfArithmeticSlices(nums []int) int {
	out, cnt := 0, 0
	for i := 2; i < len(nums); i++ {
		if nums[i]-nums[i-1] == nums[i-1]-nums[i-2] {
			cnt++
			out += cnt
		} else {
			cnt = 0
		}
	}
	return out
}
