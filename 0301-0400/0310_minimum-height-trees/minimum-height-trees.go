package main

// Tags:
// BFS

func findMinHeightTrees(n int, edges [][]int) []int {
	indeg := make([]int, n)
	next := make([][]int, n)
	for _, e := range edges {
		a, b := e[0], e[1]
		indeg[a]++
		indeg[b]++
		next[a] = append(next[a], b)
		next[b] = append(next[b], a)
	}

	q := []int{}
	for i := 0; i < n; i++ {
		if indeg[i] <= 1 {
			q = append(q, i)
		}
	}

	var out []int
	for len(q) > 0 {
		out = q
		nq := []int{}
		for _, u := range q {
			for _, v := range next[u] {
				indeg[v]--
				if indeg[v] == 1 {
					nq = append(nq, v)
				}
			}
		}
		q = nq
	}
	return out
}
