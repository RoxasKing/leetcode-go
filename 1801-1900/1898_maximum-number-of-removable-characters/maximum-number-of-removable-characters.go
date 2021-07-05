package main

// Tags:
// Binary Search

func maximumRemovals(s string, p string, removable []int) int {
	m, n := len(s), len(p)
	skip := make([]bool, len(s))
	l, r := 0, len(removable)
	for l < r {
		k := l + (r-l+1)>>1
		if isValid(s, p, m, n, removable, skip, k) {
			l = k
		} else {
			r = k - 1
		}
	}
	return l
}

func isValid(s string, p string, m, n int, removable []int, skip []bool, k int) bool {
	for i := range removable {
		skip[removable[i]] = i < k
	}
	i, j := 0, 0
	for i < m && j < n {
		if skip[i] || s[i] != p[j] {
			i++
		} else {
			i++
			j++
		}
	}
	return j == n
}
