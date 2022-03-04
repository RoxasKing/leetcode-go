package main

// Difficulty:
// Easy

// Tags:
// Dynamic Programming

func countBits(n int) []int {
	out := make([]int, n+1)
	for i := 1; i <= n; i++ {
		out[i] = out[i>>1]
		if i&1 == 1 {
			out[i]++
		}
	}
	return out
}
