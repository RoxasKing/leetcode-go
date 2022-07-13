package main

// Difficulty:
// Hard

// Tags:
// Dynamic Programming

func minCost(houses []int, cost [][]int, m int, n int, target int) int {
	f0, f1 := make([][]int, target+1), make([][]int, target+1)
	for i := 1; i <= target; i++ {
		f0[i], f1[i] = make([]int, n+1), make([]int, n+1)
		for j := 1; j <= n; j++ {
			f0[i][j], f1[i][j] = 1<<31-1, 1<<31-1
		}
	}
	f0[1][houses[0]] = 0
	if houses[0] == 0 {
		copy(f0[1][1:], cost[0])
	}
	for k := 1; k < m; k++ {
		for i := 1; i <= target; i++ {
			for j := 1; j <= n; j++ {
				if f0[i][j] == 1<<31-1 {
					continue
				}
				if houses[k] > 0 {
					if j == houses[k] {
						f1[i][j] = min(f1[i][j], f0[i][j])
					} else if i < target {
						f1[i+1][houses[k]] = min(f1[i+1][houses[k]], f0[i][j])
					}
					continue
				}
				for nj := 1; nj <= n; nj++ {
					if j == nj {
						f1[i][j] = min(f1[i][j], f0[i][j]+cost[k][nj-1])
					} else if i < target {
						f1[i+1][nj] = min(f1[i+1][nj], f0[i][j]+cost[k][nj-1])
					}
				}
			}
		}
		f0, f1 = f1, f0
		for i := 1; i <= target; i++ {
			for j := 1; j <= n; j++ {
				f1[i][j] = 1<<31 - 1
			}
		}
	}
	o := 1<<31 - 1
	for j := 1; j <= n; j++ {
		o = min(o, f0[target][j])
	}
	if o == 1<<31-1 {
		return -1
	}
	return o
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
