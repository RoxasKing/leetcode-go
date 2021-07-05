package main

// Tags:
// Binary Search
func isPerfectSquare(num int) bool {
	l, r := 0, num
	for l < r {
		m := (l + r) >> 1
		if m*m < num {
			l = m + 1
		} else {
			r = m
		}
	}
	return l*l == num
}
