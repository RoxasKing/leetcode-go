package main

// Difficulty:
// Medium

// Tags:
// Bipartite Graph
// BFS

func isBipartite(graph [][]int) bool {
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
