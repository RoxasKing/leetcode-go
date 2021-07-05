package main

import (
	"sort"
)

// Tags:
// Greedy Algorithm
func maxBuilding(n int, restrictions [][]int) int {
	m := len(restrictions)
	if m == 0 {
		return n - 1
	}

	sort.Slice(restrictions, func(i, j int) bool { return restrictions[i][0] < restrictions[j][0] })
	k := m + 1
	if restrictions[m-1][0] < n {
		k++
	}
	arr := make([][]int, k)
	copy(arr[1:1+m], restrictions)
	arr[0] = []int{1, 0}
	if k == m+2 {
		arr[k-1] = []int{n, n - 1}
	}

	for i := 1; i < k; i++ {
		h := Min(arr[i][1], arr[i-1][1]) + (arr[i][0] - arr[i-1][0])
		if h <= Max(arr[i][1], arr[i-1][1]) {
			arr[i][1] = Min(arr[i][1], h)
		}
	}
	out := arr[k-1][1]
	for i := k - 2; i >= 1; i-- {
		h := Min(arr[i][1], arr[i+1][1]) + (arr[i+1][0] - arr[i][0])
		if h <= Max(arr[i][1], arr[i+1][1]) {
			arr[i][1] = Min(arr[i][1], h)
			out = Max(out, arr[i][1])
		} else {
			out = Max(out, Max(arr[i][1], arr[i+1][1])+(arr[i+1][0]-(arr[i][0]+Abs(arr[i+1][1]-arr[i][1])))>>1)
		}
	}
	return out
}

func Abs(num int) int {
	if num < 0 {
		return -num
	}
	return num
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
