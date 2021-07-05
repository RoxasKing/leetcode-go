package main

// Tags:
// Bit Manipulation
func decode(encoded []int) []int {
	n := len(encoded) + 1
	x := 0
	for i := 1; i <= n; i++ {
		x ^= i
	}
	for i := 1; i < n-1; i += 2 {
		x ^= encoded[i]
	}
	out := make([]int, n)
	out[0] = x
	for i := 1; i < n; i++ {
		out[i] = out[i-1] ^ encoded[i-1]
	}
	return out
}
