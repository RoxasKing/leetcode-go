package main

// Difficulty:
// Hard

// Tags:
// Math
// Binary Search

func nthMagicalNumber(n int, a int, b int) int {
	c := a / gcd(a, b) * b
	l, r := 1, int(1e15)
	for l < r {
		m := l + (r-l)>>1
		if m/a+m/b-m/c < n {
			l = m + 1
		} else {
			r = m
		}
	}
	return l % (1e9 + 7)
}

func gcd(a, b int) int {
	for ; b > 0; a, b = b, a%b {
	}
	return a
}
