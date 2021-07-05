package main

import "sort"

func largestPerimeter(A []int) int {
	sort.Ints(A)
	n := len(A)
	for i := n - 3; i >= 0; i-- {
		if A[i]+A[i+1] > A[i+2] {
			return A[i] + A[i+1] + A[i+2]
		}
	}
	return 0
}
