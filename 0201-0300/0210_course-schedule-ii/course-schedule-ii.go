package main

// Tags:
// Topological Sorting + BFS
func findOrder(numCourses int, prerequisites [][]int) []int {
	edges := make([][]int, numCourses)
	indeg := make([]int, numCourses)
	for _, p := range prerequisites {
		v, u := p[0], p[1]
		edges[u] = append(edges[u], v)
		indeg[v]++
	}

	q := []int{}
	for i := 0; i < numCourses; i++ {
		if indeg[i] == 0 {
			q = append(q, i)
		}
	}

	out := make([]int, 0, numCourses)
	for len(q) > 0 {
		if len(out) == numCourses {
			return nil
		}
		u := q[0]
		q = q[1:]
		out = append(out, u)
		for _, v := range edges[u] {
			indeg[v]--
			if indeg[v] == 0 {
				q = append(q, v)
			}
		}
	}
	if len(out) < numCourses {
		return nil
	}
	return out
}
