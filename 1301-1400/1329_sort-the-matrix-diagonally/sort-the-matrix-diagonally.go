package main

import "math/rand"

// Difficulty:
// Medium

// Tags:
// Quick Sort

func diagonalSort(mat [][]int) [][]int {
	m, n := len(mat), len(mat[0])
	var qSort func(x, y, sz int)
	qSort = func(x, y, sz int) {
		if sz < 2 {
			return
		}
		pi := rand.Intn(sz)
		mat[x+sz-1][y+sz-1], mat[x+pi][y+pi] = mat[x+pi][y+pi], mat[x+sz-1][y+sz-1]
		i, j, pv := -1, sz-1, mat[x+sz-1][y+sz-1]
		for {
			for i++; i < sz-1 && mat[x+i][y+i] < pv; i++ {
			}
			for j--; j > 0 && mat[x+j][y+j] > pv; j-- {
			}
			if i >= j {
				break
			}
			mat[x+i][y+i], mat[x+j][y+j] = mat[x+j][y+j], mat[x+i][y+i]
		}
		mat[x+sz-1][y+sz-1], mat[x+i][y+i] = mat[x+i][y+i], mat[x+sz-1][y+sz-1]
		qSort(x, y, i)
		qSort(x+i+1, y+i+1, sz-i-1)
	}
	for k := 0; k < m; k++ {
		qSort(k, 0, min(m-k, n))
	}
	for k := 1; k < n; k++ {
		qSort(0, k, min(m, n-k))
	}
	return mat
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
