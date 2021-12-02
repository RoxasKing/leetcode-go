package main

// Difficulty:
// Medium

// Tags:
// DFS

func allPathsSourceTarget(graph [][]int) [][]int {
	out := [][]int{}
	dfs(graph, 0, []int{0}, &out)
	return out
}

func dfs(graph [][]int, i int, cur []int, out *[][]int) {
	if len(graph)-1 == i {
		tmp := make([]int, len(cur))
		copy(tmp, cur)
		*out = append(*out, tmp)
		return
	}
	for _, j := range graph[i] {
		cur = append(cur, j)
		dfs(graph, j, cur, out)
		cur = cur[:len(cur)-1]
	}
}
