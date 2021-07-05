package main

// Tags:
// Two Pointers

func sortArrayByParityII(nums []int) []int {
	n := len(nums)
	for i, j := 0, 1; i < n && j < n; {
		if nums[i]&1 == 1 && nums[j]&1 == 0 {
			nums[i], nums[j] = nums[j], nums[i]
			i += 2
			j += 2
			continue
		}
		if nums[i]&1 == 0 {
			i += 2
		}
		if nums[j]&1 == 1 {
			j += 2
		}
	}
	return nums
}
