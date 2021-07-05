package main

// Tags:
// Two Pointers
func maximumScore(nums []int, k int) int {
	n := len(nums)
	out := 0
	for min, l, r := nums[k], k, k; min >= 1; min-- {
		for l-1 >= 0 && nums[l-1] >= min {
			l--
		}
		for r+1 <= n-1 && nums[r+1] >= min {
			r++
		}
		out = Max(out, min*(r-l+1))
	}
	return out
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
