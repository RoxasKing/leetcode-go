package main

// Difficulty:
// Medium

// Tags:
// Binary Search

func minEatingSpeed(piles []int, h int) int {
	l, r := 1, int(1e9)
	for l < r {
		m := l + (r-l)>>1
		cnt := 0
		for _, pile := range piles {
			cnt += pile / m
			if pile%m != 0 {
				cnt++
			}
		}
		if cnt > h {
			l = m + 1
		} else {
			r = m
		}
	}
	return l
}
