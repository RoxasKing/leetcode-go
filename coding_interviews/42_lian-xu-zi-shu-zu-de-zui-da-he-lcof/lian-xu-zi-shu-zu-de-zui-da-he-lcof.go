package main

// Tags:
// Dynamic Programming
func maxSubArray(nums []int) int {
	out, sum := -1<<31, -1<<31
	for _, num := range nums {
		sum = Max(sum+num, num)
		out = Max(out, sum)
	}
	return out
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
