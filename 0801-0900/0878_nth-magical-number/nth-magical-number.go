package main

// Tags:
// Math + Binary Search
func nthMagicalNumber(n int, a int, b int) int {
	c := a / Gcd(a, b) * b
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

func Gcd(a, b int) int {
	if b == 0 {
		return a
	}
	return Gcd(b, a%b)
}
