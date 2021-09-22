package main

// Tags:
// Enumeration
// Queue
// BFS

func flipChess(chessboard []string) int {
	m, n := len(chessboard), len(chessboard[0])
	grid := make([][]byte, m)
	for i := range grid {
		grid[i] = []byte(chessboard[i])
	}
	out := 0
	for r := range grid {
		for c := range grid[r] {
			if grid[r][c] == '.' {
				solve(grid, m, n, r, c, &out)
			}
		}
	}
	return out
}

func solve(grid [][]byte, m, n, r, c int, out *int) {
	grid[r][c] = 'X'
	backup := [][]int{}

	q := [][]int{{r, c}}
	for ; len(q) > 0; q = q[1:] {
		coor := q[0]
		r, c := coor[0], coor[1]
		chg := [][]int{}
		for k := 0; k < 8; k++ {
			x, y := r, c
			i, j := dr[k], dc[k]
			cnt := 0
			tmp := [][]int{}
			for 0 <= x+i && x+i < m && 0 <= y+j && y+j < n && grid[x+i][y+j] == 'O' {
				x += i
				y += j
				cnt++
				tmp = append(tmp, []int{x, y})
			}
			if cnt > 0 && 0 <= x+i && x+i < m && 0 <= y+j && y+j < n && grid[x+i][y+j] == 'X' {
				chg = append(chg, tmp...)
			}
		}
		if len(chg) == 0 {
			continue
		}
		for _, coor := range chg {
			x, y := coor[0], coor[1]
			grid[x][y] = 'X'
		}
		q = append(q, chg...)
		backup = append(backup, chg...)
	}

	if *out < len(backup) {
		*out = len(backup)
	}
	for _, coor := range backup {
		x, y := coor[0], coor[1]
		grid[x][y] = 'O'
	}
	grid[r][c] = '.'
}

var dr = []int{0, 1, 1, 1, 0, -1, -1, -1}
var dc = []int{1, 1, 0, -1, -1, -1, 0, 1}
