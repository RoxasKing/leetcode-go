package main

// Difficulty:
// Easy

// Tags:
// Binary Search

func arrangeCoins(n int) int {
	l, r := 0, n+1
	for l < r {
		m := (l + r) >> 1
		if m*(1+m)>>1 <= n {
			l = m + 1
		} else {
			r = m
		}
	}
	return l - 1
}
