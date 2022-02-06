package main

// Difficulty:
// Medium

// Tags:
// Two Pointers

func removeDuplicates(nums []int) int {
	idx := 0
	n := len(nums)
	for i := 0; i < n; i++ {
		nums[idx] = nums[i]
		idx++
		if i+1 < n && nums[i+1] == nums[i] {
			i++
			nums[idx] = nums[i]
			idx++
		}
		for ; i+1 < n && nums[i+1] == nums[i]; i++ {
		}
	}
	return idx
}
