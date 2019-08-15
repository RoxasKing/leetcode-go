package main

import "fmt"

func minPathSum(grid [][]int) int {
	save := make([][]int, len(grid))
	for i := range save {
		save[i] = make([]int, len(grid[0]))
	}
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if i > 0 && j > 0 {
				if save[i-1][j] < save[i][j-1] {
					save[i][j] = save[i-1][j] + grid[i][j]
				} else {
					save[i][j] = save[i][j-1] + grid[i][j]
				}
			} else if i == 0 && j > 0 {
				save[i][j] = save[i][j-1] + grid[i][j]
			} else if j == 0 && i > 0 {
				save[i][j] = save[i-1][j] + grid[i][j]
			} else {
				save[i][j] = grid[i][j]
			}
		}
	}
	return save[len(grid)-1][len(grid[0])-1]
}

func main() {
	grid := [][]int{
		{1, 3, 1},
		{1, 5, 1},
		{4, 2, 1},
	}
	fmt.Println(minPathSum(grid))
}
