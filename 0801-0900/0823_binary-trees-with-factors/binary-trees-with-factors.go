package main

import "sort"

// Difficulty:
// Medium

// Tags:
// Hash
// Dynamic Programming

func numFactoredBinaryTrees(arr []int) int {
	const mod int = 1e9 + 7
	sort.Ints(arr)
	h := map[int]int{}
	for i, x := range arr {
		h[x] = i
	}
	n := len(arr)
	f := make([]int, n)
	for i := range f {
		f[i] = 1
	}
	for i := 1; i < n; i++ {
		for j := i - 1; j >= 0 && arr[j]*arr[j] >= arr[i]; j-- {
			if arr[i]%arr[j] != 0 {
				continue
			}
			if k, ok := h[arr[i]/arr[j]]; ok {
				inc := f[j] * f[k]
				if j != k {
					inc *= 2
				}
				f[i] = (f[i] + inc) % mod
			}
		}
	}
	o := 0
	for _, x := range f {
		o = (o + x) % mod
	}
	return o
}
