package main

// Tags:
// Dynamic Programming

func numberOfArithmeticSlices(nums []int) int {
	n := len(nums)
	out, cnt := 0, 0
	for i := 2; i < n; i++ {
		if nums[i]-nums[i-1] == nums[i-1]-nums[i-2] {
			cnt++
			out += cnt
		} else {
			cnt = 0
		}
	}
	return out
}
