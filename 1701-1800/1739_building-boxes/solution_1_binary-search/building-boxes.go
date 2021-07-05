package main

import "sort"

// Binary Search

func minimumBoxes(n int) int {
	l, r := 1, n
	for l < r {
		floorBoxes := l + (r-l)>>1
		if maximumBoxes(floorBoxes) < n {
			l = floorBoxes + 1
		} else {
			r = floorBoxes
		}
	}
	return l
}

func maximumBoxes(floorBoxes int) int {
	out := 0
	for floorBoxes > 0 {
		out += floorBoxes
		i := sort.Search(floorBoxes, func(i int) bool { return i*(i+1)>>1 > floorBoxes }) - 1
		next := i * (i - 1) >> 1
		if i*(i+1)>>1 < floorBoxes {
			next += floorBoxes - i*(i+1)>>1 - 1
		}
		floorBoxes = next
	}
	return out
}
