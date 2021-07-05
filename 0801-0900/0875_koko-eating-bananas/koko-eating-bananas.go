package main

// Tags:
// Binary Search
func minEatingSpeed(piles []int, h int) int {
	l, r := 1, int(1e9)
	for l < r {
		m := (l + r) >> 1
		if countHours(piles, m) > h {
			l = m + 1
		} else {
			r = m
		}
	}
	return l
}

func countHours(piles []int, limit int) int {
	out := 0
	for _, pile := range piles {
		out += pile / limit
		if pile%limit > 0 {
			out++
		}
	}
	return out
}
