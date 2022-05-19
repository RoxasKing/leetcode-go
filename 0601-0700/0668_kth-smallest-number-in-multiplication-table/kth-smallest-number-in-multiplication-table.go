package main

// Difficulty:
// Hard

// Tags:
// Binary Search

func findKthNumber(m int, n int, k int) int {
	l, r := 1, m*n
	for l < r {
		mid := l + (r-l)>>1
		c := 0
		for i, j := m, 1; i >= 1 && j <= n; {
			if i*j <= mid {
				c += i
				j++
			} else {
				i--
			}
		}
		if c < k {
			l = mid + 1
		} else {
			r = mid
		}
	}
	return l
}
