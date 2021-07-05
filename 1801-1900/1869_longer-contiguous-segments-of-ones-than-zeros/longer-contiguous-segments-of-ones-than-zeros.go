package main

// Tags:
// Two Pointers

func checkZeroOnes(s string) bool {
	n := len(s)
	max0, max1 := 0, 0
	for l, r := 0, 0; r < n; r++ {
		if s[r] != s[l] {
			l = r
		}
		if s[l] == '0' {
			max0 = Max(max0, r+1-l)
		} else {
			max1 = Max(max1, r+1-l)
		}
	}
	return max0 < max1
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
