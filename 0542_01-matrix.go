package main

/*
  给定一个由 0 和 1 组成的矩阵，找出每个元素到最近的 0 的距离。
  两个相邻元素间的距离为 1 。
*/

// BFS
func updateMatrix(matrix [][]int) [][]int {
	if len(matrix) == 0 {
		return nil
	}
	dist := make([][]int, len(matrix))
	mark := make([][]bool, len(matrix))
	for i := range dist {
		dist[i] = make([]int, len(matrix[0]))
		mark[i] = make([]bool, len(matrix[0]))
	}
	steps := [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
	var queue [][2]int
	for i := range matrix {
		for j := range matrix[0] {
			if matrix[i][j] == 0 {
				queue = append(queue, [2]int{i, j})
				mark[i][j] = true
			}
		}
	}
	for len(queue) != 0 {
		i, j := queue[0][0], queue[0][1]
		queue = queue[1:]
		for _, step := range steps {
			ni, nj := i+step[0], j+step[1]
			if 0 <= ni && ni < len(matrix) && 0 <= nj && nj < len(matrix[0]) && !mark[ni][nj] {
				dist[ni][nj] = dist[i][j] + 1
				queue = append(queue, [2]int{ni, nj})
				mark[ni][nj] = true
			}
		}
	}
	return dist
}

// Dynamic Programming
func updateMatrix2(matrix [][]int) [][]int {
	if len(matrix) == 0 {
		return nil
	}
	dist := make([][]int, len(matrix))
	for i := range dist {
		dist[i] = make([]int, len(matrix[0]))
	}
	for i := range matrix {
		for j := range matrix[0] {
			if matrix[i][j] == 0 {
				dist[i][j] = 0
			} else {
				dist[i][j] = 1<<31 - 1
			}
		}
	}
	// left & up
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[0]); j++ {
			if i-1 >= 0 {
				dist[i][j] = Min(dist[i][j], dist[i-1][j]+1)
			}
			if j-1 >= 0 {
				dist[i][j] = Min(dist[i][j], dist[i][j-1]+1)
			}
		}
	}
	// right & down
	for i := len(matrix) - 1; i >= 0; i-- {
		for j := len(matrix[0]) - 1; j >= 0; j-- {
			if i+1 < len(matrix) {
				dist[i][j] = Min(dist[i][j], dist[i+1][j]+1)
			}
			if j+1 < len(matrix[0]) {
				dist[i][j] = Min(dist[i][j], dist[i][j+1]+1)
			}
		}
	}
	return dist
}
