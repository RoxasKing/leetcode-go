package main

func rotateTheBox(box [][]byte) [][]byte {
	m, n := len(box), len(box[0])
	out := make([][]byte, n)
	for i := range out {
		out[i] = make([]byte, m)
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			out[j][m-1-i] = box[i][j]
		}
	}
	for j := 0; j < m; j++ {
		idx := n - 1
		for i := n - 1; i >= 0 && idx >= 0; i-- {
			if out[i][j] == '#' {
				out[idx][j] = '#'
				if idx != i {
					out[i][j] = '.'
				}
				idx--
			} else if out[i][j] == '*' {
				idx = i - 1
			}
		}
	}
	return out
}
