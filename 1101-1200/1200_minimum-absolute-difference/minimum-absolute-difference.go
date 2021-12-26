package main

import "sort"

// Difficulty:
// Easy

// Tags:
// Sorting

func minimumAbsDifference(arr []int) [][]int {
	sort.Ints(arr)
	out := [][]int{}
	min := 1<<31 - 1
	for i := 1; i < len(arr); i++ {
		diff := arr[i] - arr[i-1]
		if diff < min {
			min = arr[i] - arr[i-1]
			out = [][]int{{arr[i-1], arr[i]}}
		} else if diff == min {
			out = append(out, []int{arr[i-1], arr[i]})
		}
	}
	return out
}
