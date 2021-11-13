package main

// Difficulty:
// Hard

// Tags:
// Dynamic Programming

var mod int = 1e9 + 7

func kInversePairs(n int, k int) int {
	f0 := make([]int, k+1)
	f1 := make([]int, k+1)
	f0[0], f1[0] = 1, 1
	for x := 1; x <= n; x++ {
		for p := 1; p <= k && p <= x*(x-1)>>1; p++ {
			f1[p] = f1[p-1] + f0[p]
			if p-x >= 0 {
				f1[p] -= f0[p-x]
			}
			f1[p] = (f1[p] + mod) % mod
		}
		f0, f1 = f1, f0
	}
	return f0[k]
}
