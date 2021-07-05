package main

// Tags:
// BFS

func highestPeak(isWater [][]int) [][]int {
	m, n := len(isWater), len(isWater[0])
	out := make([][]int, m)
	visited := make([][]bool, m)
	for i := range out {
		out[i] = make([]int, n)
		visited[i] = make([]bool, n)
	}

	q := [][3]int{}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if isWater[i][j] == 1 {
				q = append(q, [3]int{i, j, 0})
			}
		}
	}

	for len(q) > 0 {
		e := q[0]
		q = q[1:]
		i, j, h := e[0], e[1], e[2]
		for _, f := range forwards {
			ni, nj := i+f[0], j+f[1]
			if ni < 0 || m-1 < ni || nj < 0 || n-1 < nj || isWater[ni][nj] == 1 || visited[ni][nj] {
				continue
			}
			visited[ni][nj] = true
			out[ni][nj] = h + 1
			q = append(q, [3]int{ni, nj, h + 1})
		}
	}

	return out
}

var forwards = [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
