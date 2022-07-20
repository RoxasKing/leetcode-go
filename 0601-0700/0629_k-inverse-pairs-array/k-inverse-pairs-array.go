package main

// Difficulty:
// Hard

// Tags:
// Dynamic Programming

func kInversePairs(n int, k int) int {
	const mod int = 1e9 + 7
	f0, f1 := make([]int, k+1), make([]int, k+1)
	f0[0], f1[0] = 1, 1
	for i := 1; i <= n; i++ {
		for j := 1; j <= k && j <= i*(i-1)/2; j++ {
			f1[j] = f1[j-1] + f0[j]
			if j-i >= 0 {
				f1[j] -= f0[j-i]
			}
			f1[j] = (f1[j] + mod) % mod
		}
		f0, f1 = f1, f0
	}
	return f0[k]
}
