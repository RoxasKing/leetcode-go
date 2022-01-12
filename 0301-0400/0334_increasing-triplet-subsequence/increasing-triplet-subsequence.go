package main

// Difficulty:
// Medium

// Tags:
// Greedy

func increasingTriplet(nums []int) bool {
	n := len(nums)
	if n < 3 {
		return false
	}
	for i, j, k := n-3, n-2, n-1; i >= 0; i-- {
		if nums[i] < nums[j] && nums[j] < nums[k] {
			return true
		}
		if nums[i+1] > nums[k] {
			j, k = i, i+1
		} else if nums[i] < nums[k] && nums[i] > nums[j] {
			j = i
		}
	}
	return false
}
