package main

// Tags:
// Two Pointers
func twoSum(nums []int, target int) []int {
	l, r := 0, len(nums)-1
	for l < r {
		sum := nums[l] + nums[r]
		if sum < target {
			l++
		} else if sum > target {
			r--
		} else {
			break
		}
	}
	return []int{nums[l], nums[r]}
}
