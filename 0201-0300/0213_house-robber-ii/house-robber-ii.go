package main

// Tags:
// Dynamic Programming
func rob(nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	} else if n == 1 {
		return nums[0]
	}
	return Max(dp(nums[1:]), dp(nums[:n-1]))
}

func dp(nums []int) int {
	n := len(nums)
	f := make([]int, n+1)
	f[1] = nums[0]
	for i := 2; i <= n; i++ {
		f[i] = Max(f[i-1], f[i-2]+nums[i-1])
	}
	return f[n]
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
