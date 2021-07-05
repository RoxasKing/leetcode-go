package main

func lastRemaining(n int, m int) int {
	out := 0
	for i := 2; i <= n; i++ {
		out = (out + m) % i
	}
	return out
}
