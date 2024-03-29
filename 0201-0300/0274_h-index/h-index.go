package main

import "sort"

// Sorting

func hIndex(citations []int) int {
	sort.Sort(sort.Reverse(sort.IntSlice(citations)))
	return sort.Search(len(citations), func(i int) bool { return citations[i] < i+1 })
}
