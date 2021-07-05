package main

import (
	"math/rand"
)

// Tags:
// Quick Sort + Binary Search
func bestSeqAtIndex(height []int, weight []int) int {
	var res []int
	qSortBestSeqAtIndex(height, weight, 0, len(height)-1)
	for i := range weight {
		l, r := 0, len(res)
		for l < r {
			m := l + (r-l)>>1
			if height[i] > height[res[m]] && weight[i] > weight[res[m]] {
				l = m + 1
			} else {
				r = m
			}
		}
		if l < len(res) {
			res[l] = i
		} else {
			res = append(res, i)
		}
	}
	return len(res)
}

func qSortBestSeqAtIndex(a, b []int, l, r int) {
	if l >= r {
		return
	}
	pivotIndex := l + rand.Intn(r+1-l)
	a[pivotIndex], a[r] = a[r], a[pivotIndex]
	b[pivotIndex], b[r] = b[r], b[pivotIndex]
	index := l
	for i := l; i <= r; i++ {
		if a[i] < a[r] || a[i] == a[r] && b[i] > b[r] {
			a[i], a[index] = a[index], a[i]
			b[i], b[index] = b[index], b[i]
			index++
		}
	}
	a[index], a[r] = a[r], a[index]
	b[index], b[r] = b[r], b[index]
	qSortBestSeqAtIndex(a, b, l, index-1)
	qSortBestSeqAtIndex(a, b, index+1, r)
}
