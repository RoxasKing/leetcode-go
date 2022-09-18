package main

// Difficulty:
// Medium

// Tags:
// Dynamic Programming

func maximumScore(nums []int, multipliers []int) int {
	n, m := len(nums), len(multipliers)
	f := make([][]int, m+1)
	for i := range f {
		f[i] = make([]int, m+1)
	}
	for i := 1; i <= m; i++ {
		f[i][0] = f[i-1][0] + nums[i-1]*multipliers[i-1]
		f[0][i] = f[0][i-1] + nums[n-i]*multipliers[i-1]
	}
	for l := 1; l < m; l++ {
		for r := 1; l+r <= m; r++ {
			i := l + r - 1
			f[l][r] = max(f[l-1][r]+nums[l-1]*multipliers[i], f[l][r-1]+nums[n-r]*multipliers[i])
		}
	}
	o := -1 << 31
	for i := 0; i <= m; i++ {
		o = max(o, f[i][m-i])
	}
	return o
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
