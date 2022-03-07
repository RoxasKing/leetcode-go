package main

// Difficulty:
// Hard

// Tags:
// Dynamic Programming
// Math

var mod = int(1e9 + 7)
var f = [501]int{1: 1}

func countOrders(n int) int {
	for i := 2; i <= n; i++ {
		k := (i-1)*2 + 1
		f[i] = (f[i-1] * (k + k*(k-1)/2)) % mod
	}
	return f[n]
}
