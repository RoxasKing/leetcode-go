package main

// Difficulty:
// Medium

// Tags:
// Dynamic Programming

func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	n := len(obstacleGrid[0])
	f := make([]int, n)
	f[0] = 1
	for _, os := range obstacleGrid {
		for j, o := range os {
			if o == 1 {
				f[j] = 0
			} else if j > 0 {
				f[j] += f[j-1]
			}
		}
	}
	return f[n-1]
}
