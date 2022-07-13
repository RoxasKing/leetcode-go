package main

// Difficulty:
// Medium

// Tags:
// Hash
// Dynamic Programming

func lenLongestFibSubseq(arr []int) int {
	n := len(arr)
	f := make([][]int, n)
	h := map[int]int{}
	for i, x := range arr {
		f[i] = make([]int, n)
		h[x] = i
	}
	o := 0
	for k := 2; k < n; k++ {
		for j := k - 1; j > 0 && arr[j]*2 > arr[k]; j-- {
			if i, ok := h[arr[k]-arr[j]]; ok {
				f[j][k] = max(f[i][j]+1, 3)
				o = max(o, f[j][k])
			}
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
