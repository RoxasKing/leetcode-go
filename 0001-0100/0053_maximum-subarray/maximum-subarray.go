package main

// Tags:
// Dynamic Programming
func maxSubArray(nums []int) int {
	out := -1 << 31
	cur := -1 << 31
	for _, num := range nums {
		cur = Max(cur+num, num)
		out = Max(out, cur)
	}
	return out
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
