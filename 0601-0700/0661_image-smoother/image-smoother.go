package main

// Difficulty:
// Easy

func imageSmoother(img [][]int) [][]int {
	m, n := len(img), len(img[0])
	out := make([][]int, m)
	for i := 0; i < m; i++ {
		out[i] = make([]int, n)
		for j := 0; j < n; j++ {
			sum, cnt := 0, 0
			for x := Max(0, i-1); x <= Min(m-1, i+1); x++ {
				for y := Max(0, j-1); y <= Min(n-1, j+1); y++ {
					sum += img[x][y]
					cnt++
				}
			}
			out[i][j] = sum / cnt
		}
	}
	return out
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
