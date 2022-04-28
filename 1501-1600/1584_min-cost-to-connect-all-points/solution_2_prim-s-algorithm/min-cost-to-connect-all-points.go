package main

// Difficult:
// Medium

// Tags:
// Prim's Algorithm
// Greedy

func minCostConnectPoints(points [][]int) int {
	n := len(points)
	vis := make([]bool, n)
	dis := make([]int, n)
	for i := 1; i < n; i++ {
		dis[i] = 1<<31 - 1
	}
	out := 0
	for i := 0; i < n; i++ {
		x, d := 0, 1<<31-1
		for j := 0; j < n; j++ {
			if !vis[j] && d > dis[j] {
				x, d = j, dis[j]
			}
		}
		out += d
		vis[x] = true
		for j := 0; j < n; j++ {
			if !vis[j] {
				dis[j] = Min(dis[j], Abs(points[j][0]-points[x][0])+Abs(points[j][1]-points[x][1]))
			}
		}
	}
	return out
}

func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func Abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
