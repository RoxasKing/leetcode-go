package main

// Tags:
// BFS
func floodFill(image [][]int, sr int, sc int, newColor int) [][]int {
	forwards := [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
	m, n := len(image), len(image[0])
	out := make([][]int, m)
	for i := range out {
		out[i] = make([]int, n)
		copy(out[i], image[i])
	}
	out[sr][sc] = newColor
	q := [][]int{{sr, sc}}
	for len(q) > 0 {
		e := q[0]
		q = q[1:]
		r, c := e[0], e[1]
		for _, f := range forwards {
			nr, nc := r+f[0], c+f[1]
			if nr < 0 || m-1 < nr || nc < 0 || n-1 < nc || image[nr][nc] != image[r][c] || out[nr][nc] == newColor {
				continue
			}
			out[nr][nc] = newColor
			q = append(q, []int{nr, nc})
		}
	}
	return out
}
