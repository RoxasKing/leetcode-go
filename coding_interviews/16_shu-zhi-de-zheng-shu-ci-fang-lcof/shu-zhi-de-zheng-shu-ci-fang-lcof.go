package main

func myPow(x float64, n int) float64 {
	if n < 0 {
		n = -n
		x = 1 / x
	}
	out := 1.0
	for n > 0 {
		if n&1 == 1 {
			out *= x
		}
		x *= x
		n >>= 1
	}
	return out
}
