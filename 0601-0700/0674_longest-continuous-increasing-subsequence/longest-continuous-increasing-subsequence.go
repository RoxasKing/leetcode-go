package main

func findLengthOfLCIS(nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}
	max := 1
	cur := 1
	for i := 1; i < n; i++ {
		if nums[i] > nums[i-1] {
			cur++
		} else {
			max = Max(max, cur)
			cur = 1
		}
	}
	max = Max(max, cur)
	return max
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
