package main

// Tags:
// Prefix Sum + Sliding Window
func shortestSubarray(A []int, K int) int {
	n := len(A)
	pSum := make([]int, n+1)
	for i := 0; i < n; i++ {
		pSum[i+1] = pSum[i] + A[i]
	}
	out := 1<<31 - 1

	for i, q := 0, []int{}; i <= n; i++ {
		for len(q) > 0 && pSum[i] <= pSum[q[len(q)-1]] {
			q = q[:len(q)-1]
		}
		for ; len(q) > 0 && pSum[i]-pSum[q[0]] >= K; q = q[1:] {
			out = Min(out, i-q[0])
		}
		q = append(q, i)
	}

	if out == 1<<31-1 {
		return -1
	}
	return out
}

func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
