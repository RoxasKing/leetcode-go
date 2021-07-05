package main

func trailingZeroes(n int) int {
	var out int
	for n != 0 {
		n /= 5
		out += n
	}
	return out
}
