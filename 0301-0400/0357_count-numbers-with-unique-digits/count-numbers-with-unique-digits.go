package main

// Tags:
// Math
func countNumbersWithUniqueDigits(n int) int {
	out := 1
	cur := 1
	base := 9
	for i := 1; i <= n; i++ {
		if i > 2 {
			base--
		}
		cur *= base
		out += cur
	}
	return out
}
