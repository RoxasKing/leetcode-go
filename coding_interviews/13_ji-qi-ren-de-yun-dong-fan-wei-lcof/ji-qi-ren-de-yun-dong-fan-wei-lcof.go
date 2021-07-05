package main

// Tags:
// BFS + Queue
func movingCount(m int, n int, k int) int {
	out := 0
	queue := [][]int{{0, 0}}
	visited := make([][]bool, m)
	for i := range visited {
		visited[i] = make([]bool, n)
	}
	for len(queue) > 0 {
		point := queue[0]
		queue = queue[1:]
		x, y := point[0], point[1]
		if x < 0 || m-1 < x || y < 0 || n-1 < y || visited[x][y] {
			continue
		}
		visited[x][y] = true
		if isValid(x, y, k) {
			out++
			queue = append(queue, []int{x + 1, y}, []int{x, y + 1})
		}
	}
	return out
}

func isValid(x, y, k int) bool {
	sum := 0
	for x != 0 {
		sum += x % 10
		x /= 10
	}
	for y != 0 {
		sum += y % 10
		y /= 10
	}
	return sum <= k
}

func movingCount2(m int, n int, k int) int {
	out := 1
	visited := make([][]bool, m)
	for i := range visited {
		visited[i] = make([]bool, n)
	}
	visited[0][0] = true
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if (i > 0 && visited[i-1][j] || j > 0 && visited[i][j-1]) && isValid(i, j, k) {
				out++
				visited[i][j] = true
			}
		}
	}
	return out
}
