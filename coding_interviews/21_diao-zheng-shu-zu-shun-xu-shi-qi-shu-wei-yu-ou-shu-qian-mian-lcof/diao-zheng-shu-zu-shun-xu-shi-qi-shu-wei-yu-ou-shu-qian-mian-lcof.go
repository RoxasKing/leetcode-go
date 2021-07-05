package main

// Tags:
// Two Pointers
func exchange(nums []int) []int {
	l, r := 0, len(nums)-1
	for l < r {
		if nums[l]&1 == 1 {
			l++
		} else {
			nums[l], nums[r] = nums[r], nums[l]
			r--
		}
	}
	return nums
}
