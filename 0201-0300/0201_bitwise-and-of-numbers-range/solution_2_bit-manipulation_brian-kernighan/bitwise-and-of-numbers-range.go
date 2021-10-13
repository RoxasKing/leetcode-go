package main

// Difficulty:
// Medium

// Tags:
// Bit Manipulation
// Brian Kernighan

func rangeBitwiseAnd2(m int, n int) int {
	for m < n {
		n &= (n - 1)
	}
	return n
}
