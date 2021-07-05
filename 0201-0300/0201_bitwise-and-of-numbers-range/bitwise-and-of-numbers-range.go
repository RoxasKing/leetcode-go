package main

// Tags:
// Bit Manipulation
func rangeBitwiseAnd(m int, n int) int {
	var shift int
	for m < n {
		m, n = m>>1, n>>1
		shift++
	}
	return m << shift
}

// Bit Manipulation + Brian Kernighan
func rangeBitwiseAnd2(m int, n int) int {
	for m < n {
		n &= (n - 1)
	}
	return n
}
