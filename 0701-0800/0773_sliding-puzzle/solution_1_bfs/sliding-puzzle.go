package main

// Tags:
// BFS

func slidingPuzzle(board [][]int) int {
	target := ((((1*6+2)*6+3)*6+4)*6 + 5) * 6
	visited := [6 * 6 * 6 * 6 * 6 * 6]bool{}
	visited[getIdx(board[0][0], board[0][1], board[0][2], board[1][0], board[1][1], board[1][2])] = true
	q := [][7]int{{board[0][0], board[0][1], board[0][2], board[1][0], board[1][1], board[1][2]}}
	out := -1
	for len(q) > 0 {
		e := q[0]
		q = q[1:]
		cur, t := getIdx(e[0], e[1], e[2], e[3], e[4], e[5]), e[6]
		if cur == target {
			out = t
			break
		}
		a, x0, y0 := 0, 0, 0
		for i := 0; i < 6; i++ {
			if e[i] == 0 {
				a, x0, y0 = i, i/3, i%3
				break
			}
		}
		for _, fwd := range fwds {
			x, y := x0+fwd[0], y0+fwd[1]
			if x < 0 || 1 < x || y < 0 || 2 < y {
				continue
			}
			b := x*3 + y
			ne := e
			ne[a], ne[b], ne[6] = ne[b], ne[a], t+1
			idx := getIdx(ne[0], ne[1], ne[2], ne[3], ne[4], ne[5])
			if visited[idx] {
				continue
			}
			visited[idx] = true
			q = append(q, ne)
		}
	}
	return out
}

var fwds = [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

func getIdx(a, b, c, d, e, f int) int {
	return ((((a*6+b)*6+c)*6+d)*6+e)*6 + f
}
