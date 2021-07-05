package main

import (
	"math/rand"
)

// Tags:
// Quick Sort
func getLeastNumbers(arr []int, k int) []int {
	if k == 0 {
		return nil
	}
	qSort(arr, k, 0, len(arr)-1)
	return arr[:k]
}

func qSort(arr []int, k, l, r int) {
	if l == r {
		return
	}
	pivotIdx := l + rand.Intn(r+1-l)
	arr[r], arr[pivotIdx] = arr[pivotIdx], arr[r]
	m := l
	for i := l; i < r; i++ {
		if arr[i] < arr[r] {
			arr[i], arr[m] = arr[m], arr[i]
			m++
		}
	}
	arr[m], arr[r] = arr[r], arr[m]
	if m > k-1 {
		qSort(arr, k, l, m-1)
	} else if m < k-1 {
		qSort(arr, k, m+1, r)
	}
}
