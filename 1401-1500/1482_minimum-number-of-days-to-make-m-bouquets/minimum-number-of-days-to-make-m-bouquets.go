package main

// Tags:
// Binary Search
func minDays(bloomDay []int, m int, k int) int {
	if len(bloomDay) < m*k {
		return -1
	}
	min, max := 1<<31-1, 0
	for _, flower := range bloomDay {
		min = Min(min, flower)
		max = Max(max, flower)
	}

	l, r := min, max
	for l < r {
		limit := (l + r) >> 1
		if !isValid(bloomDay, m, k, limit) {
			l = limit + 1
		} else {
			r = limit
		}
	}
	return l
}

func isValid(bloomDay []int, m, k, limit int) bool {
	cnt, cur := 0, 0
	for _, flower := range bloomDay {
		if flower > limit {
			cur = 0
		} else {
			cur++
			if cur == k {
				cnt++
				cur = 0
			}
		}
	}
	return cnt >= m
}

func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
