package main

import "sort"

// Difficulty:
// Medium

// Tags:
// Enumeration

func sequentialDigits(low int, high int) []int {
	out := []int{}
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	for l := 0; l < 9; l++ {
		for r := l; r < 9; r++ {
			v := 0
			for i := l; i <= r; i++ {
				v = v*10 + arr[i]
			}
			if low <= v && v <= high {
				out = append(out, v)
			}
		}
	}
	sort.Ints(out)
	return out
}
