package main

// Difficulty:
// Easy

// Tags:
// Two Pointers

func findErrorNums(nums []int) []int {
	for i, x := range nums {
		for j := x - 1; i != j && nums[i] != nums[j]; j = nums[i] - 1 {
			nums[i], nums[j] = nums[j], nums[i]
		}
	}
	for i, x := range nums {
		if x != i+1 {
			return []int{x, i + 1}
		}
	}
	return []int{-1, -1}
}
