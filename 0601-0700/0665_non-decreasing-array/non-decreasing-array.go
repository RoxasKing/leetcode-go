package main

// Difficulty:
// Medium

// Tags:
// Two Pointers

func checkPossibility(nums []int) bool {
	n := len(nums)
	l, r := 0, n-1
	for ; l+1 < n && nums[l] <= nums[l+1]; l++ {
	}
	for ; r-1 >= 0 && nums[r-1] <= nums[r]; r-- {
	}
	return l == n-1 || l+1 == r && ((l == 0 || nums[l-1] <= nums[r]) || (r == n-1 || nums[l] <= nums[r+1]))
}
