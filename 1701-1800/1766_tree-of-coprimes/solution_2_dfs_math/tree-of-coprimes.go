package main

// Tags:
// DFS
// Backtracking
// Math

func getCoprimes(nums []int, edges [][]int) []int {
	n := len(nums)
	es := make([][]int, n)
	for _, e := range edges {
		u, v := e[0], e[1]
		es[u] = append(es[u], v)
		es[v] = append(es[v], u)
	}

	visited := make([]bool, n)
	visited[0] = true
	parent := make([][2]int, 51)
	for i := 1; i <= 50; i++ {
		parent[i] = [2]int{-1, -1}
	}

	out := make([]int, n)
	deps := make([]int, n)
	for i := range out {
		out[i] = -1
		deps[i] = -1
	}

	dfs(n, 0, 0, nums, es, visited, parent, out, deps)

	return out
}

func dfs(n, u, dep int, nums []int, es [][]int, visited []bool, parent [][2]int, out, deps []int) {
	for i := 1; i <= 50; i++ {
		if parent[i][0] != -1 && Gcd(nums[u], i) == 1 && parent[i][1] > deps[u] {
			out[u] = parent[i][0]
			deps[u] = parent[i][1]
		}
	}
	bak := parent[nums[u]]
	parent[nums[u]] = [2]int{u, dep}
	for _, v := range es[u] {
		if visited[v] {
			continue
		}
		visited[v] = true
		dfs(n, v, dep+1, nums, es, visited, parent, out, deps)
	}
	parent[nums[u]] = bak
}

func Gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}
