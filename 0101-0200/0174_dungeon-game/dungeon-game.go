package main

// Tags:
// Dynamic Programming
func calculateMinimumHP(dungeon [][]int) int {
	m, n := len(dungeon), len(dungeon[0])
	dp := make([]int, n)
	for i := m - 1; i >= 0; i-- {
		for j := n - 1; j >= 0; j-- {
			hp := 1<<31 - 1
			if i < m-1 {
				hp = Min(hp, dp[j])
			}
			if j < n-1 {
				hp = Min(hp, dp[j+1])
			}
			if hp == 1<<31-1 {
				hp = 1
			}
			dp[j] = hp
			if hp-dungeon[i][j] > 0 {
				dp[j] = hp - dungeon[i][j]
			}
		}
	}
	return dp[0]
}

func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
