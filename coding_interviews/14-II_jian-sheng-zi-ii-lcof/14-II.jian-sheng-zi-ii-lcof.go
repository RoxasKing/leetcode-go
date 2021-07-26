package main

// Tags:
// Greedy
func cuttingRope(n int) int {
	if n <= 3 {
		return n - 1
	}
	mul3 := n / 3
	if n-mul3*3 == 1 {
		mul3--
	}
	mul2 := (n - mul3*3) >> 1
	return pow(3, mul3) * pow(2, mul2) % (1e9 + 7)
}

func pow(x, n int) int {
	out := 1
	for i := 0; i < n; i++ {
		out = (out * x) % (1e9 + 7)
	}
	return out
}
