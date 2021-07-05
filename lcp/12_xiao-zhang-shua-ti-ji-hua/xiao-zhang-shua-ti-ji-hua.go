package main

// Tags:
// Binary Search
func minTime(time []int, m int) int {
	l, r := 0, 1<<31-1
	for l < r {
		limit := (l + r) >> 1
		if !isValid(time, m, limit) {
			l = limit + 1
		} else {
			r = limit
		}
	}
	return l
}

func isValid(time []int, m, limit int) bool {
	days := 1
	sumT := 0
	maxT := -1 << 31
	for _, t := range time {
		sumT += t
		maxT = Max(maxT, t)
		if sumT-maxT > limit {
			days++
			sumT = t
			maxT = t
		}
	}
	return days <= m
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
