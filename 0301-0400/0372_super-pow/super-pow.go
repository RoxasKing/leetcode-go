package main

// Tags:
// Math

func superPow(a int, b []int) int {
	a %= mod
	out := quickMul(a, b[0])
	for _, n := range b[1:] {
		out = quickMul(out, 10) * quickMul(a, n)
		out %= mod
	}
	return out
}

func quickMul(x, n int) int {
	out := 1
	for n > 0 {
		if n&1 == 1 {
			out *= x
			out %= mod
		}
		x *= x
		x %= mod
		n >>= 1
	}
	return out
}

var mod = 1337
