package main

// Tags:
// Binary Search
func hIndex(citations []int) int {
	n := len(citations)
	l, r := 0, n
	for l < r {
		m := (l + r) >> 1
		if citations[m] < n-m {
			l = m + 1
		} else {
			r = m
		}
	}
	return n - l
}
