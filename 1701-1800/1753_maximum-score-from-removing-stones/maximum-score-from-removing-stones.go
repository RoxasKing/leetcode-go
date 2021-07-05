package main

import "sort"

// Math

func maximumScore(a int, b int, c int) int {
	arr := []int{a, b, c}
	sort.Ints(arr)
	a, b, c = arr[0], arr[1], arr[2]
	if a+b < c {
		return a + b
	}
	return c + (a+b-c)>>1
}
