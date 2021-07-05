package main

import "sort"

// Binary Search
// Prefix Sum

func findBestValue(arr []int, target int) int {
	sort.Ints(arr)
	n := len(arr)
	pSum := make([]int, n+1)
	for i := 0; i < n; i++ {
		pSum[i+1] = pSum[i] + arr[i]
	}

	l, r := 0, arr[n-1]
	for l < r {
		m := l + (r-l)>>1
		i := sort.SearchInts(arr, m)
		sum := pSum[i] + m*(n-i)
		if sum < target {
			l = m + 1
		} else {
			r = m
		}
	}

	i := sort.SearchInts(arr, l)
	j := sort.SearchInts(arr, l-1)
	if Abs((n-i)*l+pSum[i]-target) < Abs((n-j)*(l-1)+pSum[j]-target) {
		return l
	}
	return l - 1
}

func Abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}
