package main

import "sort"

// Binary Search
func findRadius(houses []int, heaters []int) int {
	m, n := len(houses), len(heaters)
	sort.Ints(houses)
	sort.Ints(heaters)
	l, r := 0, int(1e9)
	for l < r {
		radius := (l + r) >> 1
		if !isValid(houses, heaters, m, n, radius) {
			l = radius + 1
		} else {
			r = radius
		}
	}
	return l
}

func isValid(houses, heaters []int, m, n, radius int) bool {
	i, j := 0, 0
	for i < m && j < n {
		l, r := heaters[j]-radius, heaters[j]+radius
		if l <= houses[i] && houses[i] <= r {
			i++
		} else {
			j++
		}
	}
	return i == m
}
