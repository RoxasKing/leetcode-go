package main

/*
  给定一个非负整数 numRows，生成杨辉三角的前 numRows 行。
  在杨辉三角中，每个数是它左上方和右上方的数的和。
*/

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
