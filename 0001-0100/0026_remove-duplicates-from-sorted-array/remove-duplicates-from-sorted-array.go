package main

// Difficulty:
// Easy

func removeDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	o := 0
	for _, x := range nums {
		if x != nums[o] {
			o++
			nums[o] = x
		}
	}
	return o + 1
}
