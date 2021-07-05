package main

// Tags:
// Topological Sort + Dynamic Programming
func largestPathValue(colors string, edges [][]int) int {
	if len(edges) == 0 {
		return 1
	}

	n := len(colors)
	nexts := make([][]int, n)
	indeg := make([]int, n)
	for _, e := range edges {
		u, v := e[0], e[1]
		nexts[u] = append(nexts[u], v)
		indeg[v]++
	}

	f := make([][]int, n)
	for i := range f {
		f[i] = make([]int, 26)
	}

	q := []int{}
	for i := 0; i < n; i++ {
		if indeg[i] == 0 && len(nexts[i]) > 0 {
			q = append(q, i)
		}
	}

	out := -1
	for len(q) > 0 {
		u := q[0]
		q = q[1:]
		f[u][colors[u]-'a']++
		if len(nexts[u]) == 0 {
			for i := 0; i < 26; i++ {
				out = Max(out, f[u][i])
			}
		}
		for _, v := range nexts[u] {
			for i := 0; i < 26; i++ {
				f[v][i] = Max(f[v][i], f[u][i])
			}
			indeg[v]--
			if indeg[v] == 0 {
				q = append(q, v)
			}
		}
	}
	return out
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
