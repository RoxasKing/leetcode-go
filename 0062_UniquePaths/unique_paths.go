package main

import "fmt"

func uniquePaths(m int, n int) int {
	grid := make([][]int, n)
	for i := range grid {
		grid[i] = make([]int, m)
	}
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			// 到达边界的路径数都为 1
			if i == 0 || j == 0 {
				grid[i][j] = 1
				continue
			}
			// 到达 (i,j) 格子的路径数目，等于到达 上方格子 和 左边格子 路径数之和
			grid[i][j] = grid[i-1][j] + grid[i][j-1]
		}
	}
	return grid[n-1][m-1]
}

func main() {
	fmt.Println(uniquePaths(3, 7))
}
