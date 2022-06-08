package main

// Difficulty:
// Medium

// Tags:
// Binary Search

func minEatingSpeed(piles []int, h int) int {
	l, r := 1, int(1e9)
	for l < r {
		m := (l + r) / 2
		c := 0
		for _, x := range piles {
			if c += x / m; x%m > 0 {
				c++
			}
		}
		if c > h {
			l = m + 1
		} else {
			r = m
		}
	}
	return l
}
