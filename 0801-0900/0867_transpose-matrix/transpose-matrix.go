package main

// Difficulty:
// Easy

func transpose(matrix [][]int) [][]int {
	m, n := len(matrix), len(matrix[0])
	o := make([][]int, n)
	for i := range o {
		o[i] = make([]int, m)
	}
	for i, vs := range matrix {
		for j, v := range vs {
			o[j][i] = v
		}
	}
	return o
}
