package main

import "math"

// Difficulty:
// Medium

// Tags:
// Prefix Sum
// Dynamic Programming

func largestSumOfAverages(nums []int, k int) float64 {
	n := len(nums)
	pSum := make([]int, n+1)
	for i := 0; i < n; i++ {
		pSum[i+1] = pSum[i] + nums[i]
	}
	f := make([][]float64, n+1)
	for i := range f {
		f[i] = make([]float64, k+1)
	}
	for i := 1; i <= n; i++ {
		f[i][1] = float64(pSum[i]) / float64(i)
		for p := 0; p < i; p++ {
			for j := 1; j < k; j++ {
				f[i][j+1] = math.Max(f[i][j+1], f[p][j]+float64(pSum[i]-pSum[p])/float64(i-p))
			}
		}
	}
	return f[n][k]
}
