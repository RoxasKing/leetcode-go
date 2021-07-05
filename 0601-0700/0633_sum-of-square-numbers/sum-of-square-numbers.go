package main

import "math"

// Two Pointers
func judgeSquareSum(c int) bool {
	l, r := 0, int(math.Sqrt(float64(c)))
	for l <= r {
		sum := l*l + r*r
		if sum < c {
			l++
		} else if sum > c {
			r--
		} else {
			return true
		}
	}
	return false
}
