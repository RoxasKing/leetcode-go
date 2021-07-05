package main

func removeDuplicates(nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}
	idx := 0
	for i := 1; i < n; i++ {
		if nums[i] != nums[idx] {
			idx++
			nums[idx] = nums[i]
		}
	}
	nums = nums[:idx+1]
	return len(nums)
}
