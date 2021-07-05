package main

// Tags:
// Bit Manipulation
// Brian Kernighan

func hammingDistance(x int, y int) int {
	out := 0
	for xor := x ^ y; xor > 0; xor &= xor - 1 {
		out++
	}
	return out
}
