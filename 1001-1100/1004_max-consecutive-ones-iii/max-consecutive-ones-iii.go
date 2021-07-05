package main

// Tags:
// Sliding Window
func longestOnes(A []int, K int) int {
	out := 0
	l, r := 0, 0
	for r = range A {
		if A[r] == 0 {
			K--
		}
		for l <= r && K < 0 {
			if A[l] == 0 {
				K++
			}
			l++
		}
		if r+1-l > out {
			out = r + 1 - l
		}
	}
	return out
}
