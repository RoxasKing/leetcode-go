package main

// Difficulty:
// Hard

// Tags:
// Dynamic Programming

func maxCoins(nums []int) int {
	nums = append([]int{1}, append(nums, 1)...)
	n := len(nums)
	f := make([][]int, n)
	for i := range f {
		f[i] = make([]int, n)
	}
	for l := n - 3; l >= 0; l-- {
		for r := l + 2; r < n; r++ {
			for m := l + 1; m < r; m++ {
				f[l][r] = Max(f[l][r], nums[l]*nums[m]*nums[r]+f[l][m]+f[m][r])
			}
		}
	}
	return f[0][n-1]
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
