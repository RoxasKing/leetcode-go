package main

// Tags:
// Two Pointers

func sortArrayByParityII(nums []int) []int {
	i, j, n := 0, 1, len(nums)
	for i < n && j < n {
		if nums[i]&1 == 0 {
			i += 2
		} else if nums[j]&1 == 1 {
			j += 2
		} else {
			nums[i], nums[j] = nums[j], nums[i]
		}
	}
	return nums
}
