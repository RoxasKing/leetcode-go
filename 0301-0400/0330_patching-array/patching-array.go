package main

// Tags:
// Greedy
// Math

func minPatches(nums []int, n int) int {
	patches := 0
	for i, x := 0, 1; x <= n; {
		if i < len(nums) && nums[i] <= x {
			x += nums[i]
			i++
		} else {
			patches++
			x *= 2
		}
	}
	return patches
}
