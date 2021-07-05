package main

// Tags:
// Binary Search
func shipWithinDays(weights []int, D int) int {
	l, r := 0, 0
	for _, w := range weights {
		l = Max(l, w)
		r += w
	}
	for l < r {
		m := (l + r) >> 1
		if needDays(weights, m) > D {
			l = m + 1
		} else {
			r = m
		}
	}
	return l
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func needDays(weights []int, limit int) int {
	out, cur := 1, 0
	for _, w := range weights {
		if cur+w > limit {
			out++
			cur = w
		} else {
			cur += w
		}
	}
	return out
}
