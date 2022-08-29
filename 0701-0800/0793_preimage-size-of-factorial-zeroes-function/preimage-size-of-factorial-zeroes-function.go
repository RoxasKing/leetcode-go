package main

// Difficulty:
// Hard

// Tags:
// Math
// Binary Search

func preimageSizeFZF(k int) int {
	cnt0 := func(n int) int {
		o := 0
		for ; n >= 5; o += n {
			n /= 5
		}
		return o
	}
	l, r := 0, 5*k+1
	for l < r {
		m := (l + r) >> 1
		if c := cnt0(m); c == k {
			return 5
		} else if c < k {
			l = m + 1
		} else {
			r = m - 1
		}
	}
	return 0
}
