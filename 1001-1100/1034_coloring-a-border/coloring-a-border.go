package main

// Difficulty:
// Medium

// Tags:
// DFS

func colorBorder(grid [][]int, row int, col int, color int) [][]int {
	m, n, c := len(grid), len(grid[0]), grid[row][col]
	var dfs func(int, int)
	dfs = func(i, j int) {
		if i < 0 || m-1 < i || j < 0 || n-1 < j || grid[i][j] != c || grid[i][j] == 1001 {
			return
		}
		grid[i][j] = 1001
		dfs(i-1, j)
		dfs(i+1, j)
		dfs(i, j-1)
		dfs(i, j+1)
	}
	dfs(row, col)
	isTheBorder := func(i, j int) bool {
		return (i == 0 || grid[i-1][j] != 1001) ||
			(i == m-1 || grid[i+1][j] != 1001) ||
			(j == 0 || grid[i][j-1] != 1001) ||
			(j == n-1 || grid[i][j+1] != 1001)
	}
	arr1, arr2 := [][]int{}, [][]int{}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] != 1001 {
				continue
			}
			if isTheBorder(i, j) {
				arr1 = append(arr1, []int{i, j})
			} else {
				arr2 = append(arr2, []int{i, j})
			}
		}
	}
	for _, a := range arr1 {
		grid[a[0]][a[1]] = color
	}
	for _, a := range arr2 {
		grid[a[0]][a[1]] = c
	}
	return grid
}
