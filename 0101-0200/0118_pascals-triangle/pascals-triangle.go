package main

// Difficulty:
// Easy

// Tags:
// Dynamic Programming

func generate(numRows int) [][]int {
	o := make([][]int, numRows)
	for i := 0; i < numRows; i++ {
		o[i] = make([]int, i+1)
		o[i][0], o[i][i] = 1, 1
		for j := 1; j < i; j++ {
			o[i][j] = o[i-1][j-1] + o[i-1][j]
		}
	}
	return o
}
