package main

// Difficulty:
// Hard

// Tags:
// Dynamic Programming

func winnerSquareGame(n int) bool {
	f := make([]bool, n+1)
	for i := 1; i*i <= n; i++ {
		f[i*i] = true
	}
	for i := 2; i <= n; i++ {
		if f[i] {
			continue
		}
		for j := 1; j*j < i; j++ {
			if !f[i-j*j] {
				f[i] = true
				break
			}
		}
	}
	return f[n]
}
