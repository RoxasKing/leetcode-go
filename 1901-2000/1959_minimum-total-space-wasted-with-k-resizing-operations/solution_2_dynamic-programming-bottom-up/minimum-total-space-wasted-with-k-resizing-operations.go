package main

import "math"

// Tags:
// Dynamic Programming

func minSpaceWastedKResizing(nums []int, K int) int {
	N := len(nums)
	if K == N-1 {
		return 0
	}
	f := make([][]int, N)
	for i := range f {
		f[i] = make([]int, K+1)
		for k := range f[i] {
			f[i][k] = math.MaxInt32
		}
	}
	max, sum := 0, 0
	for i := 0; i < N; i++ {
		max = Max(max, nums[i])
		sum += nums[i]
		f[i][0] = max*(i+1) - sum
	}

	for i := 1; i < N; i++ {
		max, sum := 0, 0
		for j := i; j < N; j++ {
			max = Max(max, nums[j])
			sum += nums[j]
			wasted := max*(j+1-i) - sum
			for k := 1; k <= K; k++ {
				if f[i-1][k-1] == math.MaxInt32 {
					break
				}
				f[j][k] = Min(f[j][k], f[i-1][k-1]+wasted)
			}
		}
	}

	return f[N-1][K]
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
