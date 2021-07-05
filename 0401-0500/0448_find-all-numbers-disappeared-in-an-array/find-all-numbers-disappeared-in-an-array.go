package main

func findDisappearedNumbers(nums []int) []int {
	for i := range nums {
		j := nums[i] - 1
		for i != j && nums[i] != nums[j] {
			nums[i], nums[j] = nums[j], nums[i]
			j = nums[i] - 1
		}
	}
	out := []int{}
	for i := 0; i < len(nums); i++ {
		if i != nums[i]-1 {
			out = append(out, i+1)
		}
	}
	return out
}
