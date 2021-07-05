package main

// Tags:
// Dynamic Programming

func stoneGameVIII(stones []int) int {
	n := len(stones)
	pSum := make([]int, n+1)
	for i := 0; i < n; i++ {
		pSum[i+1] = pSum[i] + stones[i]
	}
	f := make([]int, n+1)
	f[n] = pSum[n]
	for i := n - 1; i >= 2; i-- {
		f[i] = Max(f[i+1], pSum[i]-f[i+1])
	}
	return f[2]
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
