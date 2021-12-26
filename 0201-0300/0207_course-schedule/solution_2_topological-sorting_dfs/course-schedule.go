package main

// Tags:
// Topological Sort + DFS
func canFinish(numCourses int, prerequisites [][]int) bool {
	edges := make([][]int, numCourses)
	visited := make([]int, numCourses) // 0: not visited 1: visiting 2: visited
	for _, p := range prerequisites {
		v, u := p[0], p[1]
		edges[u] = append(edges[u], v)
	}
	for i := 0; i < numCourses; i++ {
		if visited[i] == 0 && !dfs(edges, visited, i) {
			return false
		}
	}
	return true
}

func dfs(edges [][]int, visited []int, u int) bool {
	visited[u] = 1
	for _, v := range edges[u] {
		if visited[v] == 1 || visited[v] == 0 && !dfs(edges, visited, v) {
			return false
		}
	}
	visited[u] = 2
	return true
}
