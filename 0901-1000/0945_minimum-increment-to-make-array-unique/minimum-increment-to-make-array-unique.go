package main

import "sort"

func minIncrementForUnique(A []int) int {
	sort.Ints(A)
	out := 0
	for i := 1; i < len(A); i++ {
		if A[i] <= A[i-1] {
			out += A[i-1] + 1 - A[i]
			A[i] = A[i-1] + 1
		}
	}
	return out
}
