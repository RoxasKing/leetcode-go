package main

// Tags:
// Greedy
// Binary Search
// Math

func maxValue(n int, index int, maxSum int) int {
	maxSum -= n
	lLen := index
	rLen := n - (index + 1)
	l, r := 0, maxSum
	for l < r {
		m := (l + r) / 2
		lSum := Min(lLen, m) * (m - 1 + m - 1 - (Min(lLen, m) - 1)) / 2 // 1+2+...+Min(lLen, m-1)
		rSum := Min(rLen, m) * (m - 1 + m - 1 - (Min(rLen, m) - 1)) / 2 // 1+2+...+Min(rLen, m-1)
		if lSum+m+rSum < maxSum {
			l = m + 1
		} else {
			r = m
		}
	}

	lSum := Min(lLen, l) * (l - 1 + l - 1 - (Min(lLen, l) - 1)) / 2
	rSum := Min(rLen, l) * (l - 1 + l - 1 - (Min(rLen, l) - 1)) / 2
	if maxSum <= lSum+rSum+l-1 {
		return l
	}
	return l + 1
}

func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
