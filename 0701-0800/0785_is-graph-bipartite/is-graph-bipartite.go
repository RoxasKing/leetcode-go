package main

// Tags:
// Bipartite Graph + DFS
func isBipartite(graph [][]int) bool {
	n := len(graph)
	colors := make([]int, n)
	for i := 0; i < n; i++ {
		if colors[i] != 0 {
			continue
		}
		if ok := dfs(graph, colors, i, 1); !ok {
			return false
		}
	}
	return true
}

func dfs(graph [][]int, colors []int, node, color int) bool {
	colors[node] = color
	nextColor := 1
	if color == 1 {
		nextColor = 2
	}
	for _, neighbor := range graph[node] {
		if colors[neighbor] == 0 {
			if ok := dfs(graph, colors, neighbor, nextColor); !ok {
				return false
			}
		} else if colors[neighbor] != nextColor {
			return false
		}
	}
	return true
}

// Bipartite Graph + BFS
func isBipartite2(graph [][]int) bool {
	const UNCOLORED, RED, GREEN = 0, 1, 2
	colors := make([]int, len(graph))
	for i := range graph {
		if colors[i] == UNCOLORED {
			queue := []int{i}
			colors[i] = RED
			for len(queue) != 0 {
				node := queue[0]
				queue = queue[1:]
				nextColor := RED
				if colors[node] == RED {
					nextColor = GREEN
				}
				for _, neighbor := range graph[node] {
					if colors[neighbor] == UNCOLORED {
						queue = append(queue, neighbor)
						colors[neighbor] = nextColor
					} else if colors[neighbor] != nextColor {
						return false
					}
				}
			}
		}
	}
	return true
}
