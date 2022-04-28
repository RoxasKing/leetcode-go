package main

// Difficulty:
// Easy

// Tags:
// Two Pointers

func sortArrayByParity(nums []int) []int {
	for l, r := 0, len(nums)-1; l < r; {
		if nums[l]&1 == 0 {
			l++
		} else if nums[r]&1 == 1 {
			r--
		} else {
			nums[l], nums[r] = nums[r], nums[l]
		}
	}
	return nums
}
