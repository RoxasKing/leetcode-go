package main

// Tags:
// Dynamic Programming

func minSpaceWastedKResizing(nums []int, k int) int {
	n := len(nums)
	f := make([][]int, n)
	for i := range f {
		f[i] = make([]int, n)
		for j := range f[i] {
			f[i][j] = -1
		}
	}
	return dp(nums, f, n, k, 0)
}

func dp(nums []int, f [][]int, n, k, i int) int {
	if i == n {
		return 0
	}
	if k == -1 {
		return 1<<31 - 1
	}
	if f[i][k] != -1 {
		return f[i][k]
	}
	out := 1<<31 - 1
	max := nums[i]
	sum := 0
	for j := i; j < n; j++ {
		max = Max(max, nums[j])
		sum += nums[j]
		wasted := max*(j+1-i) - sum
		out = Min(out, dp(nums, f, n, k-1, j+1)+wasted)
	}
	f[i][k] = out
	return f[i][k]
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
