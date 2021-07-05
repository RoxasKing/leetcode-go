package main

// Tags:
// Bit Manipulation + Prefix Sum
func xorQueries(arr []int, queries [][]int) []int {
	m, n := len(arr), len(queries)
	xor := make([]int, m+1)
	for i := 0; i < m; i++ {
		xor[i+1] = xor[i] ^ arr[i]
	}
	out := make([]int, n)
	for i, q := range queries {
		l, r := q[0], q[1]
		out[i] = xor[r+1]
		if l > 0 {
			out[i] = out[i] ^ xor[l]
		}
	}
	return out
}
