package main

// Difficulty:
// Easy

// Tags:
// Sort

func missingNumber(nums []int) int {
	for i := range nums {
		for nums[i] != i && nums[i] < len(nums) {
			nums[i], nums[nums[i]] = nums[nums[i]], nums[i]
		}
	}
	for i := range nums {
		if nums[i] != i {
			return i
		}
	}
	return len(nums)
}
