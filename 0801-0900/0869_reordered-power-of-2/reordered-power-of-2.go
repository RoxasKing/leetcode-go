package main

import "sort"

// Difficulty:
// Medium

// Tags:
// Sorting

func reorderedPowerOf2(n int) bool {
	reorder := func(x int) string {
		arr := []byte{}
		for ; x != 0; x /= 10 {
			arr = append(arr, byte(x%10)+'0')
		}
		sort.Slice(arr, func(i, j int) bool { return arr[i] < arr[j] })
		return string(arr)
	}
	target := reorder(n)
	for i := 1; i <= 1e9; i <<= 1 {
		x := reorder(i)
		if x == target {
			return true
		}
	}
	return false
}
