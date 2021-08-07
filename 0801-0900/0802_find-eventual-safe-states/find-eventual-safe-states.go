package main

// Tags:
// DFS

func eventualSafeNodes(graph [][]int) []int {
	n := len(graph)
	visited := make([]bool, n)
	isSafe := make([]bool, n)
	for i := 0; i < n; i++ {
		if visited[i] {
			continue
		}
		dfs(graph, visited, isSafe, i)
	}
	var out []int
	for i := 0; i < n; i++ {
		if isSafe[i] {
			out = append(out, i)
		}
	}
	return out
}

func dfs(graph [][]int, visited, isSafe []bool, u int) {
	visited[u] = true
	for _, v := range graph[u] {
		if !visited[v] {
			dfs(graph, visited, isSafe, v)
		}
		if !isSafe[v] {
			return
		}
	}
	isSafe[u] = true
}
