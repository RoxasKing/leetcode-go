package main

/*
  给定一个非负索引 k，其中 k ≤ 33，返回杨辉三角的第 k 行。
  在杨辉三角中，每个数是它左上方和右上方的数的和。
*/

func getRow(rowIndex int) []int {
	out := make([]int, rowIndex+1)
	for i := 0; i <= rowIndex; i++ {
		out[i] = 1
		for j := i - 1; j >= 1; j-- {
			out[j] += out[j-1]
		}
	}
	return out
}
