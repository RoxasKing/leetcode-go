package main

import "sort"

// Tags:
// Greedy
// Binary Search

func canSeePersonsCount(heights []int) []int {
	n := len(heights)
	out := make([]int, n)

	arr := []int{heights[n-1]}

	for i := n - 2; i >= 0; i-- {
		j := sort.Search(len(arr), func(j int) bool { return arr[j] < heights[i] })
		out[i] = len(arr) - (j - 1)
		if j == 0 {
			out[i]--
		}
		arr = append(arr[:j], heights[i])
	}

	return out
}
