package main

// Difficulty:
// Medium

// Tags:
// Math
// Bit Manipulation

func concatenatedBinary(n int) int {
	o := 0
	for i, r, x := 1, 1, 0; i <= n; i++ {
		if i >= r {
			r <<= 1
			x++
		}
		o = (o<<x + i) % (1e9 + 7)
	}
	return o
}
