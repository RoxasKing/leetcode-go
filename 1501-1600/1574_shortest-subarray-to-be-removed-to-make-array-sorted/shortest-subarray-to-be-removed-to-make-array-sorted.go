package main

import "sort"

// Binary Search

func findLengthOfShortestSubarray(arr []int) int {
	n := len(arr)
	l, r := 0, n-1
	for l+1 <= r && arr[l] <= arr[l+1] { // 找出非递减的最大前缀 [0, l]
		l++
	}
	for l <= r-1 && arr[r-1] <= arr[r] { // 找出非递减的最大后缀 [r, n-1]
		r--
	}

	if l == r {
		return 0
	}

	out := Min(n-(l+1), r) // 比较删除 [l+1, n-1] 或 [0, r-1]，哪个删除的最少

	for i := l; i >= 0; i-- { // 当前缀为 [0, i] 时， 找出符合非递减的后缀 [j, n-1]
		j := sort.Search(n-r, func(j int) bool { return arr[j+r] >= arr[i] }) + r
		if j < n {
			out = Min(out, j-1-i)
		}
	}

	for j := r; j < n; j++ { // 当后缀为 [j, n-1] 时， 找出符合非递减的前缀 [0, i]
		i := sort.Search(l+1, func(i int) bool { return arr[i] > arr[j] }) - 1
		if i < l {
			out = Min(out, j-1-i)
		}
	}

	return out
}

func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
