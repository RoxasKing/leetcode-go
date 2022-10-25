package main

// Difficulty:
// Medium

// Tags:
// Greedy

func maxChunksToSorted(arr []int) int {
	h := [11]int{}
	for i, v := range arr {
		h[v] = i
	}
	o := 0
	for i, r := 0, 0; i < len(arr); i++ {
		if r = max(r, h[i]); i == r {
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
