package main

// Difficulty:
// Hard

// Tags:
// Dynamic Programming

func calculateMinimumHP(dungeon [][]int) int {
	m, n := len(dungeon), len(dungeon[0])
	f := make([]int, n)
	for i := m - 1; i >= 0; i-- {
		for j := n - 1; j >= 0; j-- {
			hp := 1<<31 - 1
			if i == m-1 && j == n-1 {
				hp = 1
			}
			if i < m-1 {
				hp = Min(hp, f[j])
			}
			if j < n-1 {
				hp = Min(hp, f[j+1])
			}
			if dungeon[i][j] < 0 {
				f[j] = hp - dungeon[i][j]
			} else {
				f[j] = Max(1, hp-dungeon[i][j])
			}
		}
	}
	return f[0]
}

func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
