package main

func findDiagonalOrder(mat [][]int) []int {
	m, n := len(mat), len(mat[0])
	out := make([]int, 0, m*n)
	i, j := 0, 0
	x := 1
	for len(out) < m*n {
		out = append(out, mat[i][j])
		if 0 <= i-x && i-x < m && 0 <= j+x && j+x < n {
			i += -x
			j += x
			continue
		}
		x = -x
		if i == 0 {
			if j < n-1 {
				j++
			} else {
				i++
			}
		} else {
			if i < m-1 {
				i++
			} else {
				j++
			}
		}
	}
	return out
}
