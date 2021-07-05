package main

// Tags:
// Dynamic Programming
func rob(nums []int) int {
	dp0, dp1 := 0, 0
	for _, num := range nums {
		dp0, dp1 = dp1, Max(dp1, dp0+num)
	}
	return dp1
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
