package main

// Tags:
// Bit Manipulation

func isPowerOfTwo(n int) bool {
	return n > 0 && n&(n-1) == 0
}
