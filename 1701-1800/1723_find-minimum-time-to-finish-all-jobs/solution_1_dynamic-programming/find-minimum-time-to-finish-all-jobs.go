package main

import "math/bits"

// Tags:
// Dynamic Programming

func minimumTimeRequired(jobs []int, k int) int {
	n := len(jobs)
	sum := make([]int, 1<<n)
	for i := 1; i < 1<<n; i++ {
		j := bits.TrailingZeros(uint(i))
		sum[i] = sum[i^1<<j] + jobs[j]
	}

	f0 := make([]int, 1<<n)
	f1 := make([]int, 1<<n)
	for i := 0; i < 1<<n; i++ {
		f0[i] = sum[i]
	}

	for ; k > 1; k-- {
		for i := 1; i < 1<<n; i++ {
			f1[i] = 1<<63 - 1
			for j := i; j > 0; j = (j - 1) & i {
				f1[i] = Min(f1[i], Max(f0[i-j], sum[j]))
			}
		}
		f0, f1 = f1, f0
	}

	return f0[1<<n-1]
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
