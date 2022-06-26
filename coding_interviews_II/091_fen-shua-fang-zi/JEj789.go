package main

// Difficulty:
// Medium

// Tags:
// Dynamic Programming

func minCost(costs [][]int) int {
	r, b, g := costs[0][0], costs[0][1], costs[0][2]
	for i := 1; i < len(costs); i++ {
		cost := costs[i]
		nr := cost[0] + min(b, g)
		nb := cost[1] + min(r, g)
		ng := cost[2] + min(r, b)
		r, b, g = nr, nb, ng
	}
	return min(r, min(b, g))
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
