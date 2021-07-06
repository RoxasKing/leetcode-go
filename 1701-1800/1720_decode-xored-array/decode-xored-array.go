package main

// Tags:
// Bit Manipulation

func decode(encoded []int, first int) []int {
	n := len(encoded)
	out := make([]int, n+1)
	out[0] = first
	for i := 1; i <= n; i++ {
		out[i] = out[i-1] ^ encoded[i-1]
	}
	return out
}
