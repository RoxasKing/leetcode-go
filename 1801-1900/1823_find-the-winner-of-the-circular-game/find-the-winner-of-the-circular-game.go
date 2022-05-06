package main

// Difficulty:
// Medium

// Tags:
// Math

func findTheWinner(n int, k int) int {
	last := 0
	for i := 2; i <= n; i++ {
		last = (last + k) % i
	}
	return last + 1
}
