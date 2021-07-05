package main

// Tags:
// BFS

func snakesAndLadders(board [][]int) int {
	m, n := len(board), len(board[0])
	out := -1
	q := [][2]int{{0, 0}}
	visited := [400]bool{0: true}
	for len(q) > 0 {
		e := q[0]
		q = q[1:]
		i, t := e[0], e[1]
		if i == m*n-1 {
			out = t
			break
		}
		for k := 1; k <= 6 && i+k < m*n; k++ {
			j := i + k
			x, y := j/n, j%n
			if x&1 == 1 {
				y = n - 1 - y
			}
			x = m - 1 - x
			if board[x][y] != -1 {
				j = board[x][y] - 1
			}
			if visited[j] {
				continue
			}
			visited[j] = true
			q = append(q, [2]int{j, t + 1})
		}
	}
	return out
}
