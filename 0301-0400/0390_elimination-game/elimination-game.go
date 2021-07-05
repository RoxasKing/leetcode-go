package main

// Tags:
// Math
// DFS

func lastRemaining(n int) int {
	if n == 1 {
		return 1
	}
	return 2 * (n>>1 + 1 - lastRemaining(n>>1))
}
