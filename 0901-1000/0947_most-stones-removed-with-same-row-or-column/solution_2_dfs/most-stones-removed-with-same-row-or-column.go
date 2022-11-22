package main

// Difficulty:
// Medium

// Tags:
// DFS

func removeStones(stones [][]int) int {
	stonesInRow := make(map[int][]int)
	stonesInCol := make(map[int][]int)
	for i, stone := range stones {
		stonesInRow[stone[0]] = append(stonesInRow[stone[0]], i)
		stonesInCol[stone[1]] = append(stonesInCol[stone[1]], i)
	}
	visitedRow := make(map[int]bool)
	visitedCol := make(map[int]bool)
	n := len(stones)
	visited := make([]bool, n)
	count := 0
	for i := n - 1; i >= 0; i-- {
		if visitedRow[stones[i][0]] || visitedCol[stones[i][1]] || visited[i] {
			continue
		}
		count++
		dfs(stones, stonesInRow, stonesInCol, visitedRow, visitedCol, visited, i)
	}
	return n - count
}

func dfs(
	stones [][]int,
	stonesInRow, stonesInCol map[int][]int,
	visitedRow, visitedCol map[int]bool,
	visited []bool, i int,
) {
	if !visitedRow[stones[i][0]] {
		visitedRow[stones[i][0]] = true
		for _, stoneIdx := range stonesInRow[stones[i][0]] {
			if !visited[stoneIdx] {
				visited[stoneIdx] = true
				dfs(stones, stonesInRow, stonesInCol, visitedRow, visitedCol, visited, stoneIdx)
			}
		}
	}
	if !visitedCol[stones[i][1]] {
		visitedCol[stones[i][1]] = true
		for _, stoneIdx := range stonesInCol[stones[i][1]] {
			if !visited[stoneIdx] {
				visited[stoneIdx] = true
				dfs(stones, stonesInRow, stonesInCol, visitedRow, visitedCol, visited, stoneIdx)
			}
		}
	}
}
