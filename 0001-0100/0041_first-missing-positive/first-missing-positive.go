package main

func firstMissingPositive(nums []int) int {
	n := len(nums)
	for i := range nums {
		j := nums[i] - 1
		for 0 <= j && j < n && i != j && nums[j] != nums[i] {
			nums[j], nums[i] = nums[i], nums[j]
			j = nums[i] - 1
		}
	}

	for i := range nums {
		if i+1 != nums[i] {
			return i + 1
		}
	}
	return n + 1
}
