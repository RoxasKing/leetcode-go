package main

func findErrorNums(nums []int) []int {
	for i := range nums {
		for j := nums[i] - 1; i != j && nums[i] != nums[j]; j = nums[i] - 1 {
			nums[i], nums[j] = nums[j], nums[i]
		}
	}

	var out []int
	for i := range nums {
		if i+1 != nums[i] {
			out = []int{nums[i], i + 1}
			break
		}
	}
	return out
}
