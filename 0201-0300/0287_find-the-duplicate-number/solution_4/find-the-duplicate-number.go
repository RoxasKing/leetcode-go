package main

// Difficulty:
// Medium

func findDuplicate(nums []int) int {
	for i := range nums {
		for j := nums[i] - 1; i != j+1 && nums[i] != nums[j]; j = nums[i] - 1 {
			nums[i], nums[j] = nums[j], nums[i]
		}
	}
	for i, num := range nums {
		if i+1 != num {
			return num
		}
	}
	return len(nums)
}
