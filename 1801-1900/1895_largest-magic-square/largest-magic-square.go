package main

// Tags:
// Prefix Sum

func largestMagicSquare(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	f := make([][]int, m+1)
	f[0] = make([]int, n+1)
	for i := 0; i < m; i++ {
		f[i+1] = make([]int, n+1)
		for j := 0; j < n; j++ {
			f[i+1][j+1] = grid[i][j] + f[i][j+1] + f[i+1][j] - f[i][j]
		}
	}

	out := 1
	for x := 0; x < m; x++ {
		for y := 0; y < n; y++ {
			for k := out + 1; x+k <= m && y+k <= n; k++ {
				if valid(grid, f, x, y, k) {
					out = k
				}
			}
		}
	}
	return out
}

func valid(grid, f [][]int, x, y, k int) bool {
	sum := f[x+k][y+k] - f[x+k][y] - f[x][y+k] + f[x][y]
	if sum%k > 0 {
		return false
	}

	target := sum / k
	d1, d2 := 0, 0

	for i := 0; i < k; i++ {
		if f[x+i+1][y+k]-f[x+i+1][y]-f[x+i][y+k]+f[x+i][y] != target {
			return false
		}
		if f[x+k][y+i+1]-f[x+k][y+i]-f[x][y+i+1]+f[x][y+i] != target {
			return false
		}
		d1 += grid[x+i][y+i]
		d2 += grid[x+k-1-i][y+i]
	}

	return d1 == target && d2 == target
}
