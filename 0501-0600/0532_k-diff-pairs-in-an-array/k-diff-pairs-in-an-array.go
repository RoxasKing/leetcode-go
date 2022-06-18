package main

// Difficulty:
// Medium

// Tags:
// Hash

func findPairs(nums []int, k int) int {
	h := map[int]int{}
	for _, x := range nums {
		h[x]++
	}
	o := 0
	for x, v := range h {
		if k == 0 && v > 1 || k > 0 && h[x+k] > 0 {
			o++
		}
	}
	return o
}
