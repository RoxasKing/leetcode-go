package main

import "sort"

// Difficulty:
// Easy

// Tags:
// Sorting

func canBeEqual(target []int, arr []int) bool {
	sort.Ints(target)
	sort.Ints(arr)
	for i := range arr {
		if target[i] != arr[i] {
			return false
		}
	}
	return true
}
