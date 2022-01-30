package main

// Difficulty:
// Medium

// Tags:
// BFS

func highestPeak(isWater [][]int) [][]int {
	m, n := len(isWater), len(isWater[0])
	out := make([][]int, m)
	for i := range out {
		out[i] = make([]int, n)
		for j := range out[i] {
			out[i][j] = -1
		}
	}
	q := [][]int{}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if isWater[i][j] == 1 {
				out[i][j] = 0
				continue
			}
			if i-1 >= 0 && isWater[i-1][j] == 1 ||
				i+1 < m && isWater[i+1][j] == 1 ||
				j-1 >= 0 && isWater[i][j-1] == 1 ||
				j+1 < n && isWater[i][j+1] == 1 {
				out[i][j] = 1
				q = append(q, []int{i, j})
			}
		}
	}
	dr := []int{-1, 0, 1, 0}
	dc := []int{0, 1, 0, -1}
	for ; len(q) > 0; q = q[1:] {
		cell := q[0]
		i, j := cell[0], cell[1]
		height := out[i][j] + 1
		for k := 0; k < 4; k++ {
			r, c := i+dr[k], j+dc[k]
			if r < 0 || m-1 < r || c < 0 || n-1 < c {
				continue
			}
			if out[r][c] == -1 || height < out[r][c] {
				out[r][c] = height
				q = append(q, []int{r, c})
			}
		}
	}
	return out
}
