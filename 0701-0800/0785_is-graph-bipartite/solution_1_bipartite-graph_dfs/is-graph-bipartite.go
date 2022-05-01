package main

// Difficulty:
// Medium

// Tags:
// Bipartite Graph
// DFS

func isBipartite(graph [][]int) bool {
	n := len(graph)
	colors := make([]int, n)
	var dfs func(int, int) bool
	dfs = func(u, color int) bool {
		colors[u] = color
		nextColor := 1
		if color == 1 {
			nextColor = 2
		}
		for _, v := range graph[u] {
			if colors[v] == nextColor {
				continue
			}
			if colors[v] != 0 {
				return false
			}
			if ok := dfs(v, nextColor); !ok {
				return false
			}
		}
		return true
	}
	for i := range colors {
		if colors[i] != 0 {
			continue
		}
		if ok := dfs(i, 1); !ok {
			return false
		}
	}
	return true
}
