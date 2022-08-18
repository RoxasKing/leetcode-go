package main

// Difficulty:
// Hard

// Tags:
// Greedy

func maxChunksToSorted(arr []int) int {
	n := len(arr)
	maxl := make([]int, n)
	minr := make([]int, n)
	maxl[0] = arr[0]
	minr[n-1] = arr[n-1]
	for i := 1; i < n; i++ {
		maxl[i] = max(maxl[i-1], arr[i])
		minr[n-1-i] = min(minr[n-i], arr[n-1-i])
	}
	o := 1
	for i := 0; i < n-1; i++ {
		if maxl[i] <= minr[i+1] {
			o++
		}
	}
	return o
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
