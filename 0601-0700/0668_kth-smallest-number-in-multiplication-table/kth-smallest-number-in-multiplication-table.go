package main

// Difficulty:
// Hard

// Tags:
// Binary Search

func findKthNumber(m int, n int, k int) int {
	l, r := 1, m*n
	for l < r {
		limit := (l + r) >> 1
		if countNoBigger(m, n, limit) < k {
			l = limit + 1
		} else {
			r = limit
		}
	}
	return l
}

func countNoBigger(m, n, limit int) int {
	out := 0
	row, col := m, 1
	for row >= 1 && col <= n {
		if row*col <= limit {
			out += row
			col++
		} else {
			row--
		}
	}
	return out
}
