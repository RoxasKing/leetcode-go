package main

func printNumbers(n int) []int {
	size := 1
	for i := 0; i < n; i++ {
		size *= 10
	}
	out := make([]int, size-1)
	for i := range out {
		out[i] = i + 1
	}
	return out
}
