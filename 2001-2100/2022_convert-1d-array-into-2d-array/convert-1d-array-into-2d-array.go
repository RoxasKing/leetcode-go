package main

// Difficulty:
// Easy

func construct2DArray(original []int, m int, n int) [][]int {
	if len(original) != m*n {
		return nil
	}
	out := make([][]int, m)
	for i := range out {
		out[i] = make([]int, n)
	}
	for i, num := range original {
		out[i/n][i%n] = num
	}
	return out
}
