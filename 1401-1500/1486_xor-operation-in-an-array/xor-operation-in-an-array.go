package main

// Tags:
// Bit Manipulation
func xorOperation(n int, start int) int {
	out := start
	for i := 1; i < n; i++ {
		out ^= start + 2*i
	}
	return out
}
