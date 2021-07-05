package main

// Tags:
// Math + Binary Search

func nthUglyNumber(n int, a int, b int, c int) int {
	d, e, f := a/Gcd(a, b)*b, b/Gcd(b, c)*c, a/Gcd(a, c)*c
	g := a / Gcd(a, e) * e
	l, r := 1, 4*int(1e9)+1
	for l < r {
		m := (l + r) >> 1
		if m/a+m/b+m/c-m/d-m/e-m/f+m/g < n {
			l = m + 1
		} else {
			r = m
		}
	}
	return l
}

func Gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}
