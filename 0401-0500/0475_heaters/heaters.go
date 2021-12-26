package main

import "sort"

// Difficulty:
// Medium

// Tags:
// Binary Search
// Two Pointers

func findRadius(houses []int, heaters []int) int {
	sort.Ints(houses)
	sort.Ints(heaters)
	l, r := 0, int(1e9)
	for l < r {
		radius := (l + r) >> 1
		if !isValid(houses, heaters, radius) {
			l = radius + 1
		} else {
			r = radius
		}
	}
	return l
}

func isValid(houses, heaters []int, radius int) bool {
	i, j := 0, 0
	for i < len(houses) && j < len(heaters) {
		if Abs(houses[i]-heaters[j]) <= radius {
			i++
		} else {
			j++
		}
	}
	return i == len(houses)
}

func Abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
