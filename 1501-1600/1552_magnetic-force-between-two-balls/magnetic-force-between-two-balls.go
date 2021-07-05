package main

import "sort"

// Binary Search

func maxDistance(position []int, m int) int {
	sort.Ints(position)
	n := len(position)
	l, r := 1, position[n-1]
	for l < r {
		force := (l + r + 1) >> 1
		if canPutBalls(position, n, force) >= m {
			l = force
		} else {
			r = force - 1
		}
	}
	return l
}

func canPutBalls(position []int, n, force int) int {
	out := 1
	pre := position[0]
	for i := 1; i < n; i++ {
		if pre+force <= position[i] {
			out++
			pre = position[i]
		}
	}
	return out
}
