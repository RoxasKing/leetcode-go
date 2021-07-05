package main

func numWays(n int) int {
	a, b := 0, 1
	for i := 1; i <= n; i++ {
		a, b = b, (a+b)%(1e9+7)
	}
	return b
}
