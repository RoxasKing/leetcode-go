package leetcode

/*
  给定一个非负整数 numRows，生成杨辉三角的前 numRows 行。
  在杨辉三角中，每个数是它左上方和右上方的数的和。
*/

func generate(numRows int) [][]int {
	index := -1
	out := make([][]int, numRows)
	for i := 0; i < numRows; i++ {
		tmp := make([]int, i+1)
		tmp[0], tmp[i] = 1, 1
		for j := 1; j < i; j++ {
			tmp[j] = out[index][j-1] + out[index][j]
		}
		index++
		out[index] = tmp
	}
	return out
}
