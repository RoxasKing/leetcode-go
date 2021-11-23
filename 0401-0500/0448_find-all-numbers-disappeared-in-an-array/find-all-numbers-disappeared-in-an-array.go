package main

// Difficulty:
// Easy

func findDisappearedNumbers(nums []int) []int {
	n := len(nums)
	for i := 0; i < n; i++ {
		for j := nums[i] - 1; j <= n && i != j && nums[i] != nums[j]; j = nums[i] - 1 {
			nums[i], nums[j] = nums[j], nums[i]
		}
	}
	out := []int{}
	for i := 0; i < n; i++ {
		if i+1 != nums[i] {
			out = append(out, i+1)
		}
	}
	return out
}
