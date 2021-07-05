package main

func generate(numRows int) [][]int {
	out := make([][]int, numRows)
	for i := 0; i < numRows; i++ {
		out[i] = make([]int, i+1)
		out[i][0], out[i][i] = 1, 1
		for j := 1; j < i; j++ {
			out[i][j] = out[i-1][j-1] + out[i-1][j]
		}
	}
	return out
}
